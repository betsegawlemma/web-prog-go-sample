package entity

import "time"

// Category represents Food Menu Category
type Category struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
	Image       string `gorm:"type:varchar(255)"`
	Items       []Item `gorm:"many2many:item_categories"`
}

// User represents application user
type User struct {
	ID       uint
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null; unique"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
	Orders   []Order
}

// Role repesents application user roles
type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}

// Item represents food menu items
type Item struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Price       float32
	Description string
	Categories  []Category   `gorm:"many2many:item_categories"`
	Image       string       `gorm:"type:varchar(255)"`
	Ingredients []Ingredient `gorm:"many2many:item_ingredients"`
}

// Ingredient represents ingredients in a food item
type Ingredient struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
}

// Order represents customer order
type Order struct {
	ID        uint
	CreatedAt time.Time
	UserID    uint
	ItemID    uint
	Quantity  uint
}

//Session represents login user session
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

// Comment represents comments forwarded by application users
type Comment struct {
	ID        uint
	FullName  string `gorm:"type:varchar(255)"`
	Message   string
	Phone     string `gorm:"type:varchar(100);not null; unique"`
	Email     string `gorm:"type:varchar(255);not null; unique"`
	CreatedAt time.Time
}
