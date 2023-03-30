package DataBase

import (
	"fmt"

	"github.com/manoj771vsj/Routines/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToCarDB() *gorm.DB {
	DB_URL := "host=localhost user=Arpita password=brocode dbname=library port=5432 sslmode=disable" //Dtabse URl
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to Database")
	}
	fmt.Print("Successfully connected to Database")
	db.AutoMigrate(&entity.Car{})

	DB := db
	return DB
}

func ConnectOwnerDB() *gorm.DB {
	DB_URL := "host=localhost user=Arpita password=brocode dbname=library port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to Database")
	}
	fmt.Print("Successfully connected to Database")
	db.AutoMigrate(&entity.Owner{})
	DB := db
	return DB
}

func ConnectCarDetailsDB() *gorm.DB {
	DB_URL := "host=localhost user=Arpita password=brocode dbname=library port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to Database")
	}
	fmt.Print("Successfully connected to Database")
	db.AutoMigrate(&entity.CarDetails{})
	DB := db
	return DB
}
