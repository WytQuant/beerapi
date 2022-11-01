package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"komgrip-api/config"
	"komgrip-api/routes"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Fail to load .env file")
	}

	// connect to database
	config.InitMYSQL()
	config.InitMONGO()

	// Create server
	server := gin.Default()

	server.Static("/uploads", "./uploads")

	uploadDirs := []string{"beers"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}

	routes.Serve(server)

	port := "4000"

	if err := server.Run(":" + port); err != nil {
		log.Fatalln("Fail to listen server port:", port)
	}
}
