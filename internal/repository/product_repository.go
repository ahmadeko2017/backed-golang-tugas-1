package repository

import (
	"github.com/ahmadeko2017/backed-golang-tugas-1/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *entity.Product) error
	FindAll() ([]entity.Product, error)
	FindByID(id uint) (entity.Product, error)
	Update(product *entity.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	// Preload Category to include category details in response
	err := r.db.Preload("Category").Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (entity.Product, error) {
	var product entity.Product
	err := r.db.Preload("Category").First(&product, id).Error
	return product, err
}

func (r *productRepository) Update(product *entity.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}
