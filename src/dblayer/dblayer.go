package dblayer

import "neulhan-commerce-server/src/models"

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProductByID(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(email, pass string) (models.Customer, error)
	SignOutUserByID(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}
