package handler

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS
// func SetCORS(r *gin.Engine) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		r.Use(cors.New(cors.Config{
// 			AllowOrigins:     []string{"http://localhost:3000"},
// 			AllowMethods:     []string{"*"},
// 			AllowHeaders:     []string{"*"},
// 			AllowCredentials: true,
// 			MaxAge:           24 * time.Hour,
// 		}))
// 	}
// }
func SetCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
}

// func SetCORS() gin.HandlerFunc {
// 	config := cors.DefaultConfig()
// 	config.AllowOrigins = []string{"http://localhost:3000"}
// 	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
// 	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
// 	config.AllowCredentials = true
// 	config.MaxAge = 24 * time.Hour

// 	return cors.New(config)
// }
