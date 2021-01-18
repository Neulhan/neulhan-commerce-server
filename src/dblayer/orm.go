package dblayer

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"neulhan-commerce-server/src/models"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbName string, con *gorm.Config) (*DBORM, error) {
	db, err := gorm.Open(postgres.Open(dbName), con)
	if err != nil {
		log.Fatal("DATABASE OPEN FAILED")
	}
	err = db.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{})
	if err != nil {
		log.Fatal("DATABASE MIGRATE FAILED")
	}
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetCustomerByName(name string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{Name: name}).Find(customer).Error
}

func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(customer, id).Error
}

func (db *DBORM) GetProductByID(id uint) (product models.Product, err error) {
	return product, db.First(product, id).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	return customer, db.Create(&customer).Error
}

func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	if !checkPassword(pass) {
		return customer, errors.New("InvalidPassword")
	}
	result := db.Table("Customer").Where(&models.Customer{Email: email})

	err = result.Update("logged_in", 1).Error
	if err != nil {
		return customer, err
	}
	return customer, result.Find(&customer).Error
}

func (db *DBORM) SignOutUserByID(id int) error {
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	return db.Table("Customer").Where(&customer).Update("logged_id", 0).Error
}

func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("Order").Select("*").Joins("join customer on customer.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id=?", id).Scan(&orders).Error
}

func checkPassword(pass string) (check bool) {
	return check
}

func hashPassword(pass *string) (hashedPass string) {
	return hashedPass
}
