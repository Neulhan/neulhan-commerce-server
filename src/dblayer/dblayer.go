package dblayer

import "neulhan-commerce-server/src/models"

type DBLayer interface {
	GetAllProduct() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser() (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserByID(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}
