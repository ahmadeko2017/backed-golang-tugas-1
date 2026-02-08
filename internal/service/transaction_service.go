package service

import (
	"errors"
	"time"

	"github.com/ahmadeko2017/backed-golang-tugas/internal/entity"
	"github.com/ahmadeko2017/backed-golang-tugas/internal/repository"
	"github.com/ahmadeko2017/backed-golang-tugas/pkg/database"
	"gorm.io/gorm"
)

type CheckoutItem struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type TransactionService interface {
	Checkout(items []CheckoutItem, total float64) (*entity.Transaction, error)
	ReportRange(start time.Time, end time.Time) (float64, int64, string, int64, error)
	ReportToday() (float64, int64, string, int64, error)
}

type transactionService struct {
	txRepo   repository.TransactionRepository
	prodRepo repository.ProductRepository
	db       *gorm.DB
}

func NewTransactionService(txRepo repository.TransactionRepository, prodRepo repository.ProductRepository) TransactionService {
	return &transactionService{txRepo: txRepo, prodRepo: prodRepo, db: database.DB}
}

func (s *transactionService) Checkout(items []CheckoutItem, total float64) (*entity.Transaction, error) {
	if len(items) == 0 {
		return nil, errors.New("no items provided")
	}

	// Run atomic transaction
	var resultTx *entity.Transaction
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var computedTotal float64

		// Validate products and stock first
		products := make(map[uint]entity.Product)
		for _, it := range items {
			if it.Quantity <= 0 {
				return errors.New("invalid quantity for product")
			}
			// Use Lock to prevent race condition
			p, err := s.prodRepo.FindByIDWithLock(tx, it.ProductID)
			if err != nil {
				return errors.New("product not found")
			}
			if p.Stock < it.Quantity {
				return errors.New("insufficient stock for product: " + p.Name)
			}
			products[it.ProductID] = p
			computedTotal += p.Price * float64(it.Quantity)
		}

		if computedTotal != total {
			return errors.New("total price mismatch")
		}

		// Create transaction record
		txRecord := &entity.Transaction{Total: total}
		if err := s.txRepo.Create(tx, txRecord); err != nil {
			return err
		}

		// Create details and update stocks
		for _, it := range items {
			p := products[it.ProductID]
			detail := entity.TransactionDetail{
				TransactionID: txRecord.ID,
				ProductID:     p.ID,
				Quantity:      it.Quantity,
				Price:         p.Price,
			}
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}

			// decrement stock
			p.Stock = p.Stock - it.Quantity
			if p.Stock < 0 {
				return errors.New("stock would become negative")
			}
			if err := tx.Save(&p).Error; err != nil {
				return err
			}
		}

		resultTx = txRecord
		return nil
	})

	return resultTx, err
}

func (s *transactionService) ReportRange(start time.Time, end time.Time) (float64, int64, string, int64, error) {
	return s.txRepo.ReportRange(start, end)
}

func (s *transactionService) ReportToday() (float64, int64, string, int64, error) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := start.Add(24*time.Hour - time.Nanosecond)
	return s.ReportRange(start, end)
}
