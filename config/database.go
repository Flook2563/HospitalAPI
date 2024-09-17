package config

import (
	"fmt"
	"log"

	"github.com/Flook2563/Hospitalapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "host=localhost user=hospital_user password=hospital_password dbname=hospital_db port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect Database Fail ! :", err)
	} else {
		fmt.Println("Connect Database Success !")
	}

	err = DB.AutoMigrate(&models.Staff{}, &models.Hospital{}, &models.Patient{})
	if err != nil {
		log.Fatal("AutoMigrate Fail ! :", err)
	} else {
		fmt.Println("AutoMigrate Success !")
	}
}
