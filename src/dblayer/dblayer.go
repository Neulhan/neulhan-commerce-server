package dblayer

import "neulhan-commerce-server/src/models"

type DBLayer interface {
	GetProducts() ([]models.Product, error)
	GetProduct(id int) (models.Product, error)
	CreateProduct(models.Product) (models.Product, error)
	UpdateProduct(models.Product) (models.Product, error)
	DeleteProduct(int) error
	GetProductByID(uint) (models.Product, error)
	GetPromos() ([]models.Product, error)
	GetUserByName(string) (models.User, error)
	GetUserByID(int) (models.User, error)
	GetUsers() ([]models.User, error)
	AddUser(models.User) (models.User, error)
	SignInUser(email, pass string) (models.User, error)
	SignOutUserByID(int) error
	DeleteUserByID(int) error
	GetUserOrdersByID(int) ([]models.Order, error)
}
