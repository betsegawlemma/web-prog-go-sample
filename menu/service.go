package menu

import "github.com/betsegawlemma/webproggob/entity"

// CategoryService specifies food menu category related services
type CategoryService interface {
	Categories() ([]entity.Category, error)
	StoreCategores(categoies []entity.Category) error
}
