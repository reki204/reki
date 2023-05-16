package handler

import "github.com/gin-gonic/gin"

// クライアントへの返すレスポンスの型
type Responses struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// エラーのレスポンスの型
type ErrRes struct {
	StatusCode int    `json:"statusCode"`
	Title      string `json:"title"`
	Message    string `json:"message"`
}

// APIレスポンスの作成（成功時）
func APIResponse(ctx *gin.Context, StatusCode int, Message string, Data interface{}) {
	jsonResponse := &Responses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}
	ctx.JSON(StatusCode, jsonResponse)
}

// エラーレスポンス作成
func ErrResponse(ctx *gin.Context, StatusCode int, Title, Message string) {
	res := ErrRes{
		StatusCode: StatusCode,
		Title:      Title,
		Message:    Message,
	}
	defer ctx.AbortWithStatus(StatusCode)
	ctx.JSON(StatusCode, res)
}
