package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
	Orders   []Order
	Cart     Cart
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryID  uint
}

type Category struct {
	gorm.Model
	Name        string
	Description string
	Products    []Product
}

type Order struct {
	gorm.Model
	UserID  uint
	Status  string // pending, paid, shipped
	Total   float64
	Items   []OrderItem
	Payment Payment
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
}

type Cart struct {
	gorm.Model
	UserID uint
	Items  []CartItem
}

type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Quantity  int
}

type Payment struct {
	gorm.Model
	OrderID uint
	Method  string // cod, vnpay, etc.
	PaidAt  *time.Time
	Status  string // pending, completed, failed
}
