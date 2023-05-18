package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reki204/reki/handler"
)

func SetRouting(ctx context.Context, router *gin.Engine) error {
	router.NoRoute(handleNotFound)
	router.NoMethod(handleInternalServerError)

	// ルーティング設定
	groupRoute := router.Group("/api/v1")
	apiCheckhandler := handler.NewApiCheckHandler()
	groupRoute.GET("/", apiCheckhandler.GetMessage)

	return nil
}

// 404 Not Found handler
func handleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}

// 500 Internal Server Error handler
func handleInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
	})
}
