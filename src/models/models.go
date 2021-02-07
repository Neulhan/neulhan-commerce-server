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

type User struct {
	gorm.Model
	Name     string  `json:"name" gorm:"column: name"`
	Email    string  `json:"email" gorm:"column:email"`
	SocialID string  `json:"socialID"`
	Social   string  `json:"social"`
	LoggedIn bool    `json:"loggedIn" gorm:"column:logged_in"`
	Orders   []Order `json:"orders"`
}

func (User) TableName() string {
	return "users"
}

type Order struct {
	gorm.Model
	Product
	User
	UserID       int       `json:"userID" gorm:"column:user_id"`
	ProductID    int       `json:"productID" gorm:"column:product_id"`
	Price        float64   `json:"sellPrice" gorm:"column:price"`
	PurchaseDate time.Time `json:"purchaseDate" gorm:"column:purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}

type Session struct {
	gorm.Model
	UserID int `json:"userID" gorm:"column:user_id"`
}

//func (s *Session) GetUser() User {
//}
