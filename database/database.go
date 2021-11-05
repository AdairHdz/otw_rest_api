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
			migrateTables()
			seedTables()
		}
	}

	return db, err
}

func migrateTables() {
	if db == nil {
		return
	}
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

func seedTables() {
	if db == nil {
		return
	}

	states := []entity.State{
		{
			EntityUUID: entity.EntityUUID{
				ID: "395bd679-79b4-4517-9c27-72a92db797b3",
			},
			Name: "Veracruz",
			Cities: []entity.City{
				{
					EntityUUID: entity.EntityUUID{
						ID: "4e22c768-6e2c-4626-94a9-099b3732f9ac",
					},
					Name: "Xalapa",
				},
				{
					EntityUUID: entity.EntityUUID{
						ID: "2360ce3c-6c26-45af-8f53-93b96405c6ae",
					},
					Name: "Coatepec",
				},
				{
					EntityUUID: entity.EntityUUID{
						ID: "ef89c7a8-a6b1-41bb-ad6a-e677dc66ebee",
					},
					Name: "Xico",
				},
				{
					EntityUUID: entity.EntityUUID{
						ID: "160b4fc4-6ca4-4b35-95c4-6714db0381b0",
					},
					Name: "Tomatlán",
				},
				{
					EntityUUID: entity.EntityUUID{
						ID: "9cc86989-44c0-42df-b39b-3f450f15d5b9",
					},
					Name: "Teocelo",
				},
				{
					EntityUUID: entity.EntityUUID{
						ID: "e2692316-3985-4acc-8825-122fbcd03b9b",
					},
					Name: "Minatitlán",
				},
				{
					EntityUUID: entity.EntityUUID{
						ID: "57c1b25d-2041-44af-aa17-e9681976457d",
					},
					Name: "Córdoba",
				},
			},
		},
	}

	db.Save(&states)
	workingDays := []entity.WorkingDay{
		{
			ID: 1,
			Name: "Lunes",
		},
		{
			ID: 2,
			Name: "Martes",
		},
		{
			ID: 3,
			Name: "Miércoles",
		},
		{
			ID: 4,
			Name: "Jueves",
		},
		{
			ID: 5,
			Name: "Viernes",
		},
		{
			ID: 6,
			Name: "Sábado",
		},
		{
			ID: 7,
			Name: "Domingo",
		},
	}
	db.Save(&workingDays)
}