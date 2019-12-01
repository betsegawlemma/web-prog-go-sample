package service

import (
	"github.com/betsegawlemma/restaurant/entity"
	"github.com/betsegawlemma/restaurant/menu"
)

// CategoryServiceImpl implements menu.CategoryService interface
type CategoryServiceImpl struct {
	categoryRepo menu.CategoryRepository
}

// NewCategoryServiceImpl will create new CategoryService object
func NewCategoryServiceImpl(CatRepo menu.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{categoryRepo: CatRepo}
}

// Categories returns list of categories
func (cs *CategoryServiceImpl) Categories() ([]entity.Category, error) {

	categories, err := cs.categoryRepo.Categories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// StoreCategory persists new category information
func (cs *CategoryServiceImpl) StoreCategory(category entity.Category) error {

	err := cs.categoryRepo.StoreCategory(category)

	if err != nil {
		return err
	}

	return nil
}

// Category returns a category object with a given id
func (cs *CategoryServiceImpl) Category(id int) (entity.Category, error) {

	c, err := cs.categoryRepo.Category(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateCategory updates a cateogory with new data
func (cs *CategoryServiceImpl) UpdateCategory(category entity.Category) error {

	err := cs.categoryRepo.UpdateCategory(category)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory delete a category by its id
func (cs *CategoryServiceImpl) DeleteCategory(id int) error {

	err := cs.categoryRepo.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}
