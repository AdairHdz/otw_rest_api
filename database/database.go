package database

import (	
	"fmt"
	"os"
	"github.com/AdairHdz/OTW-Rest-API/entity"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	databaseUser string
	databasePassword string	
	databaseName string
	databaseContainerName string
)

func New() (*gorm.DB, error) {
	databaseUser = os.Getenv("DB_USER")
	databasePassword = os.Getenv("DB_PASSWORD")	
	databaseName = os.Getenv("DB_NAME")
	databaseContainerName = os.Getenv("DB_CONTAINER_NAME")

	var err error
	if db == nil {						
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			databaseUser, databasePassword, databaseContainerName, databaseName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err == nil {
			db.AutoMigrate(&entity.State{})
			db.AutoMigrate(&entity.City{})
			db.AutoMigrate(&entity.Account{})
			db.AutoMigrate(&entity.User{})
			db.AutoMigrate(&entity.ServiceRequester{})
			db.AutoMigrate(&entity.ServiceProvider{})			
			db.AutoMigrate(&entity.Address{})
			db.AutoMigrate(&entity.PriceRate{})
			db.AutoMigrate(&entity.Review{})
			db.AutoMigrate(&entity.Score{})
			db.AutoMigrate(&entity.ServiceRequest{})
			db.AutoMigrate(&entity.Evidence{})
			db.AutoMigrate(&entity.WorkingDay{})
		}		
	}

	return db, err
}