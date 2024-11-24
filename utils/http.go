package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccessJson(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func ResponseFailJson(ctx *gin.Context, reason string) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"reason":  reason,
	})
}
