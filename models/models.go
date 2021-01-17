package models

import "time"

type Product struct {
	Image       string  `json:"img"`
	ImageAlt    string  `json:"imgAlt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	ProductName string  `json:"productName"`
	Description string  `json:"desc"`
}

type Customer struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	LoggedIn bool   `json:"loggedIn"`
}

type Order struct {
	Product
	Customer
	CustomerID   int       `json:"customerID"`
	ProductID    int       `json:"productID"`
	Price        float64   `json:"sellPrice"`
	PurchaseDate time.Time `json:"purchaseDate"`
}
