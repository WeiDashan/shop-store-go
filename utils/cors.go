package utils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		//AllowOrigins:     []string{"https://foo.com"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Accept", "token"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//MaxAge:           12 * time.Hour,
	})
}
