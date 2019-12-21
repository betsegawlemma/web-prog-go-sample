package menu

import "github.com/betsegawlemma/restaurant/entity"

// CategoryRepository specifies food menu category database operations
type CategoryRepository interface {
	Categories() ([]entity.Category, []error)
	Category(id uint) (*entity.Category, []error)
	UpdateCategory(category *entity.Category) (*entity.Category, []error)
	DeleteCategory(id uint) (*entity.Category, []error)
	StoreCategory(category *entity.Category) (*entity.Category, []error)
	ItemsInCategory(category *entity.Category) ([]entity.Item, []error)
}

// ItemRepository specifies food menu item related database operations
type ItemRepository interface {
	Items() ([]entity.Item, []error)
	Item(id uint) (*entity.Item, []error)
	UpdateItem(menu *entity.Item) (*entity.Item, []error)
	DeleteItem(id uint) (*entity.Item, []error)
	StoreItem(item *entity.Item) (*entity.Item, []error)
}

// RoleRepository speifies application user role related database operations
type RoleRepository interface {
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

// IngredientRepository speifies food item ingredints related database operations
type IngredientRepository interface {
	Ingredients() ([]entity.Ingredient, []error)
	Ingredient(id uint) (*entity.Ingredient, []error)
	UpdateIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error)
	DeleteIngredient(id uint) (*entity.Ingredient, []error)
	StoreIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error)
}

// UserRepository specifies application user related database operations
type UserRepository interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
}

// OrderRepository specifies customer menu order related database operations
type OrderRepository interface {
	Orders() ([]entity.Order, []error)
	Order(id uint) (*entity.Order, []error)
	CustomerOrders(customer *entity.User) ([]entity.Order, []error)
	UpdateOrder(order *entity.Order) (*entity.Order, []error)
	DeleteOrder(id uint) (*entity.Order, []error)
	StoreOrder(order *entity.Order) (*entity.Order, []error)
}

// CommentRepository specifies customer comment related database operations
type CommentRepository interface {
	Comments() ([]entity.Comment, []error)
	Comment(id uint) (*entity.Comment, []error)
	UpdateComment(comment *entity.Comment) (*entity.Comment, []error)
	DeleteComment(id uint) (*entity.Comment, []error)
	StoreComment(comment *entity.Comment) (*entity.Comment, []error)
}
