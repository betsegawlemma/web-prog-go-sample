package repository

import (
	"errors"

	"github.com/betsegawlemma/restaurant/entity"
	"github.com/betsegawlemma/restaurant/menu"
	"github.com/jinzhu/gorm"
)

// MockCategoryRepo implements the menu.CategoryRepository interface
type MockCategoryRepo struct {
	conn *gorm.DB
}

// NewMockCategoryRepo will create a new object of MockCategoryRepo
func NewMockCategoryRepo(db *gorm.DB) menu.CategoryRepository {
	return &MockCategoryRepo{conn: db}
}

// Categories returns all fake categories
func (mCatRepo *MockCategoryRepo) Categories() ([]entity.Category, []error) {
	ctgs := []entity.Category{entity.CategoryMock}
	return ctgs, nil
}

// Category retrieve a fake category with id 1
func (mCatRepo *MockCategoryRepo) Category(id uint) (*entity.Category, []error) {
	ctg := entity.CategoryMock
	if id == 1 {
		return &ctg, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateCategory updates a given fake category
func (mCatRepo *MockCategoryRepo) UpdateCategory(category *entity.Category) (*entity.Category, []error) {
	cat := entity.CategoryMock
	return &cat, nil
}

// DeleteCategory deletes a given category from the database
func (mCatRepo *MockCategoryRepo) DeleteCategory(id uint) (*entity.Category, []error) {
	cat := entity.CategoryMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &cat, nil
}

// StoreCategory stores a given mock category
func (mCatRepo *MockCategoryRepo) StoreCategory(category *entity.Category) (*entity.Category, []error) {
	cat := category
	return cat, nil
}

// ItemsInCategory returns mock food menu items
func (mCatRepo *MockCategoryRepo) ItemsInCategory(category *entity.Category) ([]entity.Item, []error) {
	items := []entity.Item{entity.ItemMock}
	return items, nil
}
