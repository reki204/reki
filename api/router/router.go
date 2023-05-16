package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reki204/reki/handler"
)

func SetRouting(ctx context.Context, router *gin.Engine) error {
	// 404エラー
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	})

	// 500エラー
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	})

	// ルーティング設定
	groupRoute := router.Group("/api/v1")
	apiCheckhandler := handler.NewApiCheckHandler()
	groupRoute.GET("/", apiCheckhandler.GetMessage)

	return nil
}
