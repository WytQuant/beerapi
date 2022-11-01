package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"komgrip-api/models"
	"log"
	"os"
)

var db *gorm.DB
var collection *mongo.Collection

const (
	DBNAME  = "loggerAPI"
	COLNAME = "logInfo"
)

func InitMYSQL() {
	var err error
	db, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalln("Fail to connect to database")
	}

	if err := db.AutoMigrate(&models.Beer{}); err != nil {
		log.Fatalln("Fail to migration with given model")
	}
}

func GetMYSQL() *gorm.DB {
	return db
}

func InitMONGO() {
	mongoURI := os.Getenv("MONGO_URI")
	clientOption := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatalln(err.Error())
	}

	collection = client.Database(DBNAME).Collection(COLNAME)
}

func GetMONGO() *mongo.Collection {
	return collection
}
