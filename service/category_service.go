package service

import (
	"errors"
	"project3/model/entity"
	"project3/model/input"
	"project3/repository"
)

type CategoryService interface {
	CreateCategory(role_user string, input input.CategoryCreateInput) (entity.Category, error)
	GetAllCategory() ([]entity.Category, error)
	GetTasksByCategoryID(id_category int) ([]entity.Task, error)
	PatchCategory(role_user string, id_category int, input input.CategoryPatchInput) (entity.Category, error)
	DeleteCategory(role_user string, id_category int) (entity.Category, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryService {
	return &categoryService{categoryRepository}
}

func (s *categoryService) CreateCategory(role_user string, input input.CategoryCreateInput) (entity.Category, error) {
	if role_user != "admin" {
		return entity.Category{}, errors.New("you are not admin")
	}

	category := entity.Category{
		Type: input.Type,
	}

	categoryData, err := s.categoryRepository.CreateCategory(category)
	if err != nil {
		return entity.Category{}, err
	}

	return categoryData, nil
}
func (s *categoryService) GetAllCategory() ([]entity.Category, error) {
	categories, err := s.categoryRepository.GetAll()
	if err != nil {
		return []entity.Category{}, err
	}

	return categories, nil
}
func (s *categoryService) GetTasksByCategoryID(id_category int) ([]entity.Task, error) {
	tasks, err := s.categoryRepository.GetTasksByCategoryID(id_category)
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (s *categoryService) PatchCategory(role_user string, id_category int, input input.CategoryPatchInput) (entity.Category, error) {
	if role_user != "admin" {
		return entity.Category{}, errors.New("you are not admin")
	}

	category := entity.Category{
		Type: input.Type,
	}

	_, err := s.categoryRepository.PatchType(id_category, category)
	if err != nil {
		return entity.Category{}, err
	}

	categoryData, err := s.categoryRepository.GetCategoryById(id_category)
	if err != nil {
		return entity.Category{}, err
	}

	return categoryData, nil
}
func (s *categoryService) DeleteCategory(role_user string, id_category int) (entity.Category, error) {
	if role_user != "admin" {
		return entity.Category{}, errors.New("you are not admin")
	}

	categoryData, err := s.categoryRepository.GetCategoryById(id_category)
	if err != nil {
		return entity.Category{}, err
	}
	if categoryData.ID == 0 {
		return entity.Category{}, errors.New("data not found")
	}

	categoryData, err = s.categoryRepository.Delete(id_category)
	if err != nil {
		return entity.Category{}, err
	}

	return categoryData, nil
}
