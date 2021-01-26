package dblayer

import (
	"gorm.io/gorm"
	"neulhan-commerce-server/src/models"
)

func (db *DBORM) GetProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) CreateProduct(product models.Product) (models.Product, error) {
	return product, db.Create(&product).Error
}

func (db *DBORM) UpdateProduct(updateProduct models.Product) (product models.Product, err error) {
	return updateProduct, db.Table("Product").Where(models.Product{
		Model: gorm.Model{
			ID: updateProduct.ID,
		},
	}).Save(&updateProduct).Error
}

func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetProductByID(id uint) (product models.Product, err error) {
	return product, db.First(product, id).Error
}
