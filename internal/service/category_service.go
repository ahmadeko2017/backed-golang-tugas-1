package service

import (
	"github.com/ahmadeko2017/backed-golang-tugas/internal/entity"
	"github.com/ahmadeko2017/backed-golang-tugas/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *entity.Category) error
	GetAllCategories() ([]entity.Category, error)
	GetCategoryByID(id uint) (entity.Category, error)
	UpdateCategory(id uint, category *entity.Category) error
	DeleteCategory(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) CreateCategory(category *entity.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) GetAllCategories() ([]entity.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetCategoryByID(id uint) (entity.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) UpdateCategory(id uint, input *entity.Category) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	category.Name = input.Name
	category.Description = input.Description

	return s.repo.Update(&category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}
