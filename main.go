package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"komgrip-api/config"
	"komgrip-api/controllers"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Fail to load .env file")
	}

	// connect to database
	config.InitDB()
	db := config.GetDB()

	// Create server
	server := gin.Default()

	server.Static("/uploads", "./uploads")

	uploadDirs := []string{"beers"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}

	beerController := controllers.BeerController{DB: db}
	beerGroup := server.Group("/api/v1")
	{
		beerGroup.GET("/beer", beerController.GetAll)
		beerGroup.POST("/beer", beerController.Create)
		beerGroup.PUT("beer/:id", beerController.Update)
		beerGroup.DELETE("/beer/:id", beerController.Delete)
	}

	port := "4000"

	if err := server.Run(":" + port); err != nil {
		log.Fatalln("Fail to listen server port:", port)
	}
}
