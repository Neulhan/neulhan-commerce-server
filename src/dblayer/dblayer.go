package dblayer

import "neulhan-commerce-server/src/models"

type DBLayer interface {
	GetAllProduct() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(models.Customer) (models.Customer, error)
	SignOutUserByID(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}
