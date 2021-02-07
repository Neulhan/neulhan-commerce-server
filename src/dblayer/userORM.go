package dblayer

import (
	"errors"
	"gorm.io/gorm"
	"neulhan-commerce-server/src/models"
)

func (db *DBORM) GetUserByName(name string) (user models.User, err error) {
	return user, db.Where(&models.User{Name: name}).Find(user).Error
}

func (db *DBORM) GetUserByID(id int) (user models.User, err error) {
	return user, db.First(user, id).Error
}

func (db *DBORM) AddUser(user models.User) (models.User, error) {
	hashPassword(&user.Pass)
	user.LoggedIn = true
	return user, db.Create(&user).Error
}

func (db *DBORM) SignInUser(email, pass string) (user models.User, err error) {
	if !checkPassword(pass) {
		return user, errors.New("InvalidPassword")
	}
	result := db.Table("User").Where(&models.User{Email: email})

	err = result.Update("logged_in", 1).Error
	if err != nil {
		return user, err
	}
	return user, result.Find(&user).Error
}

func (db *DBORM) SignOutUserByID(id int) error {
	user := models.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	return db.Table("User").Where(&user).Update("logged_id", 0).Error
}

func (db *DBORM) GetUsers() (users []models.User, err error) {
	return users, db.Find(&users).Error
}

func (db *DBORM) GetUserOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("Order").Select("*").Joins("join user on user.id = user_id").Joins("join products on products.id = product_id").Where("user_id=?", id).Scan(&orders).Error
}

func checkPassword(pass string) (check bool) {
	return check
}

func hashPassword(pass *string) (hashedPass string) {
	return hashedPass
}
