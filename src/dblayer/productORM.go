package dblayer

import (
	"neulhan-commerce-server/src/models"
)

func (db *DBORM) GetProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) CreateProduct(product models.Product) (models.Product, error) {
	return product, db.Create(&product).Error
}

func (db *DBORM) UpdateProduct(updateProduct models.Product) (product models.Product, err error) {
	return updateProduct, db.First(&product, updateProduct.ID).Updates(updateProduct).Error
}

func (db *DBORM) DeleteProduct(productToDelete models.Product) error {
	return db.Delete(&productToDelete).Error
}

func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetProductByID(id uint) (product models.Product, err error) {
	return product, db.First(product, id).Error
}
