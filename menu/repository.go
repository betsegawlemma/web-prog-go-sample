package menu

import "github.com/betsegawlemma/restaurant/entity"

// CategoryRepository specifies menu category related database operations
type CategoryRepository interface {
	Categories() ([]entity.Category, error)
	Category(id int) (entity.Category, error)
	UpdateCategory(category entity.Category) error
	DeleteCategory(id int) error
	StoreCategory(category entity.Category) error
}
