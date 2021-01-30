package dblayer

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"neulhan-commerce-server/src/models"
	"os"
	"time"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbName string, con *gorm.Config) (*DBORM, error) {
	var err error
	var db *gorm.DB
	for {
		db, err = gorm.Open(postgres.Open(dbName), con)
		if err != nil {
			fmt.Println(dbName)
			fmt.Printf("FAILED -> RECONNECT TO DATABASE[%s]", os.Getenv("MODE"))

			time.Sleep(time.Second * 3)
			continue
		}
		err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
		if err != nil {
			log.Fatal("DATABASE MIGRATE FAILED")
		}
		break
	}

	fmt.Println("DATABASE CONNECTED!")
	return &DBORM{
		DB: db,
	}, err
}
