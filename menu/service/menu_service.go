package service

import (
	"github.com/betsegawlemma/restaurant/entity"
	"github.com/betsegawlemma/restaurant/menu"
)

// CategoryService implements menu.CategoryService interface
type CategoryService struct {
	categoryRepo menu.CategoryRepository
}

// NewCategoryService will create new CategoryService object
func NewCategoryService(CatRepo menu.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: CatRepo}
}

// Categories returns list of categories
func (cs *CategoryService) Categories() ([]entity.Category, error) {

	categories, err := cs.categoryRepo.Categories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// StoreCategory persists new category information
func (cs *CategoryService) StoreCategory(category entity.Category) error {

	err := cs.categoryRepo.StoreCategory(category)

	if err != nil {
		return err
	}

	return nil
}

// Category returns a category object with a given id
func (cs *CategoryService) Category(id int) (entity.Category, error) {

	c, err := cs.categoryRepo.Category(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateCategory updates a cateogory with new data
func (cs *CategoryService) UpdateCategory(category entity.Category) error {

	err := cs.categoryRepo.UpdateCategory(category)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory delete a category by its id
func (cs *CategoryService) DeleteCategory(id int) error {

	err := cs.categoryRepo.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}
