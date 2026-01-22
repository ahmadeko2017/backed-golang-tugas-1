package repository

import (
	"github.com/ahmadeko2017/backed-golang-tugas-1/internal/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *entity.Category) error
	FindAll() ([]entity.Category, error)
	FindByID(id uint) (entity.Category, error)
	Update(category *entity.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(category *entity.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindByID(id uint) (entity.Category, error) {
	var category entity.Category
	err := r.db.First(&category, id).Error
	return category, err
}

func (r *categoryRepository) Update(category *entity.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Category{}, id).Error
}
