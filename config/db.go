package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"komgrip-api/models"
	"log"
	"os"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalln("Fail to connect to database")
	}

	if err := db.AutoMigrate(&models.Beer{}); err != nil {
		log.Fatalln("Fail to migration with given model")
	}
}

func GetDB() *gorm.DB {
	return db
}
