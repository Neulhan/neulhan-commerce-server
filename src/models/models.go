package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImageAlt    string  `json:"imgAlt" gorm:"column:img_alt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	ProductName string  `json:"productName" gorm:"column:product_name"`
	Description string  `json:"desc"`
}

func (Product) TableName() string {
	return "products"
}

type Customer struct {
	gorm.Model
	Name     string  `json:"name" gorm:"column: name"`
	Email    string  `json:"email" gorm:"column:email"`
	Pass     string  `json:"password"`
	LoggedIn bool    `json:"loggedIn" gorm:"column:logged_in"`
	Orders   []Order `json:"orders"`
}

func (Customer) TableName() string {
	return "customers"
}

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `json:"customerID" gorm:"column:customer_id"`
	ProductID    int       `json:"productID" gorm:"column:product_id"`
	Price        float64   `json:"sellPrice" gorm:"column:price"`
	PurchaseDate time.Time `json:"purchaseDate" gorm:"column:purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}
