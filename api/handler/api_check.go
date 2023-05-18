package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiCheck struct{}

func NewApiCheckHandler() *ApiCheck {
	return &ApiCheck{}
}

func (ac *ApiCheck) GetMessage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}
