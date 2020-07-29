package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Badmsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 500,
		"info":   msg,
	})
}

func Goodmsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   msg,
	})
}
