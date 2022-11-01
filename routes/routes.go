package routes

import (
	"github.com/gin-gonic/gin"
	"komgrip-api/config"
	"komgrip-api/controllers"
	"komgrip-api/middlewares"
)

func Serve(r *gin.Engine) {
	db := config.GetMYSQL()
	beerController := controllers.BeerController{DB: db}
	beerGroup := r.Group("/api/v1")
	beerGroup.GET("/beer", beerController.GetAll)
	beerGroup.Use(middlewares.LoggingInfoMiddleware())
	{
		beerGroup.POST("/beer", beerController.Create)
		beerGroup.PUT("beer/:id", beerController.Update)
		beerGroup.DELETE("/beer/:id", beerController.Delete)
	}
}
