package menu

import "github.com/betsegawlemma/restaurant/entity"

// CategoryService specifies food menu category services
type CategoryService interface {
	Categories() ([]entity.Category, []error)
	Category(id uint) (*entity.Category, []error)
	UpdateCategory(category *entity.Category) (*entity.Category, []error)
	DeleteCategory(id uint) (*entity.Category, []error)
	StoreCategory(category *entity.Category) (*entity.Category, []error)
	ItemsInCategory(category *entity.Category) ([]entity.Item, []error)
}

// ItemService specifies food menu item related services
type ItemService interface {
	Items() ([]entity.Item, []error)
	Item(id uint) (*entity.Item, []error)
	UpdateItem(menu *entity.Item) (*entity.Item, []error)
	DeleteItem(id uint) (*entity.Item, []error)
	StoreItem(item *entity.Item) (*entity.Item, []error)
}

// RoleService speifies application user role related services
type RoleService interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

// IngredientService speifies food item ingredints related services
type IngredientService interface {
	Ingredients() ([]entity.Ingredient, []error)
	Ingredient(id uint) (*entity.Ingredient, []error)
	UpdateIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error)
	DeleteIngredient(id uint) (*entity.Ingredient, []error)
	StoreIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error)
}

// UserService specifies application user related services
type UserService interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
}

//OrderService specifies customer menu order related services
type OrderService interface {
	Orders() ([]entity.Order, []error)
	Order(id uint) (*entity.Order, []error)
	CustomerOrders(customer *entity.User) ([]entity.Order, []error)
	UpdateOrder(order *entity.Order) (*entity.Order, []error)
	DeleteOrder(id uint) (*entity.Order, []error)
	StoreOrder(order *entity.Order) (*entity.Order, []error)
}

// CommentService specifies customer comment related service
type CommentService interface {
	Comments() ([]entity.Comment, []error)
	Comment(id uint) (*entity.Comment, []error)
	UpdateComment(comment *entity.Comment) (*entity.Comment, []error)
	DeleteComment(id uint) (*entity.Comment, []error)
	StoreComment(comment *entity.Comment) (*entity.Comment, []error)
}
