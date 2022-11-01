package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"komgrip-api/config"
	"komgrip-api/models"
	"log"
	"time"
)

func LoggingInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().Format(time.RFC1123)
		c.Next()

		collection := config.GetMONGO()
		logInfo := &models.LogInfo{
			RequestAt:  start,
			StatusCode: c.Writer.Status(),
			Path:       c.Request.RequestURI,
			Method:     c.Request.Method,
			ClientAddr: c.ClientIP(),
		}

		_, err := collection.InsertOne(context.Background(), logInfo)
		if err != nil {
			log.Fatalln(err.Error())
		}

	}
}
