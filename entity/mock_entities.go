package entity

import "time"

// CategoryMock mocks Food Menu Category
var CategoryMock = Category{
	ID:          1,
	Name:        "Mock Category 01",
	Description: "Mock Category 01 Description",
	Image:       "mock_cat.png",
	Items:       []Item{},
}

// RoleMock mocks user role entity
var RoleMock = Role{
	ID:    1,
	Name:  "Mock Role 01",
	Users: []User{},
}

// ItemMock mocks food menu items
var ItemMock = Item{
	ID:          1,
	Name:        "Mock Item 01",
	Price:       50.5,
	Description: "Mock Item 01 Description",
	Categories:  []Category{},
	Image:       "mock_item.png",
	Ingredients: []Ingredient{},
}

// IngredientMock mocks ingredients in a food item
var IngredientMock = Ingredient{
	ID:          1,
	Name:        "Mock Ingredient 01",
	Description: "Mock Ingredient 01 Description",
}

// OrderMock mocks customer order
var OrderMock = Order{
	ID:        1,
	CreatedAt: time.Time{},
	UserID:    1,
	ItemID:    1,
	Quantity:  100,
}

// UserMock mocks application user
var UserMock = User{
	ID:       1,
	FullName: "Mock User 01",
	Email:    "mockuser@example.com",
	Phone:    "0900000000",
	Password: "P@$$w0rd",
	RoleID:   1,
	Orders:   []Order{},
}

// SessionMock mocks sessions of loged in user
var SessionMock = Session{
	ID:         1,
	UUID:       "_session_one",
	SigningKey: []byte("RestaurantApp"),
	Expires:    0,
}

// CommentMock mocks comments forwarded by application users
var CommentMock = Comment{
	ID:        1,
	FullName:  "Mock User 01",
	Message:   "Mock message",
	Phone:     "0900000000",
	Email:     "mockuser@example.com",
	CreatedAt: time.Time{},
}
