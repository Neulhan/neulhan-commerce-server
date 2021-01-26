package dblayer

import (
	"neulhan-commerce-server/src/models"
)

func (db *dblayer.DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *dblayer.DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *dblayer.DBORM) GetProductByID(id uint) (product models.Product, err error) {
	return product, db.First(product, id).Error
}
