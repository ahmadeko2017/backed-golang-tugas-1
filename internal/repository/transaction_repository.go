package repository

import (
	"time"

	"github.com/ahmadeko2017/backed-golang-tugas/internal/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *gorm.DB, transaction *entity.Transaction) error
	ReportRange(start time.Time, end time.Time) (float64, int64, string, int64, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(tx *gorm.DB, transaction *entity.Transaction) error {
	return tx.Create(transaction).Error
}

// ReportRange returns total revenue, total transactions, top product name and qty
func (r *transactionRepository) ReportRange(start time.Time, end time.Time) (float64, int64, string, int64, error) {
	var totalRevenue float64
	var totalTx int64
	// sum totals
	if err := r.db.Model(&entity.Transaction{}).
		Where("created_at >= ? AND created_at <= ?", start, end).
		Select("COALESCE(SUM(total),0)").Scan(&totalRevenue).Error; err != nil {
		return 0, 0, "", 0, err
	}
	if err := r.db.Model(&entity.Transaction{}).
		Where("created_at >= ? AND created_at <= ?", start, end).
		Count(&totalTx).Error; err != nil {
		return 0, 0, "", 0, err
	}

	// top product
	type topRow struct {
		Name string
		Qty  int64
	}
	var row topRow
	// join transaction_details -> products, sum qty group by product
	if err := r.db.Table("transaction_details").
		Select("products.name as name, COALESCE(SUM(transaction_details.quantity),0) as qty").
		Joins("left join products on products.id = transaction_details.product_id").
		Joins("left join transactions on transactions.id = transaction_details.transaction_id").
		Where("transactions.created_at >= ? AND transactions.created_at <= ?", start, end).
		Group("products.name").
		Order("qty desc").
		Limit(1).
		Scan(&row).Error; err != nil {
		return totalRevenue, totalTx, "", 0, err
	}

	if row.Name == "" || row.Qty == 0 {
		return totalRevenue, totalTx, "", 0, nil
	}
	return totalRevenue, totalTx, row.Name, row.Qty, nil
}
