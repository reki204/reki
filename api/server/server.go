package server

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	if err := r.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}

func router() *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	r.Use(setCORS(r))

	// ルーティング
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})

	// 404エラー
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
	})

	// 500エラー
	r.NoMethod(func(c *gin.Context) {
		c.JSON(500, gin.H{
				"message": "Internal Server Error",
		})
	})

	return r
}

// CORS
func setCORS(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		r.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:3000",},
			AllowMethods: []string{"POST","GET","OPTIONS",},
			AllowHeaders: []string{"Content-Type",},
			AllowCredentials: true,
			MaxAge: 24 * time.Hour,
		}))
		c.Next()
	}
}
