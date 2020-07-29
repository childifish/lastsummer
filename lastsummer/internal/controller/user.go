package controller

import (
	"github.com/gin-gonic/gin"
	"lastsummer/internal/middleware"
	"lastsummer/internal/service"
	"lastsummer/resp"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		resp.Badmsg(ctx, "Binding error")
		return
	}

	ok, msg := service.Register(user.Username, user.Password)
	if !ok {
		resp.Badmsg(ctx, msg)
		return
	}

	resp.Goodmsg(ctx, "注册成功，请重新登录"+"localhost:8080/user/login")
}

func Login(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		resp.Badmsg(ctx, "Binding error")
		return
	}

	ok := service.Login(user.Username, user.Password)
	if !ok {
		resp.Badmsg(ctx, "用户名或密码错误")
		return
	}

	token, err := middleware.GenerateToken(user.Username, user.Password)
	if err != nil {
		log.Print(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 10000, "TOKEN": token})

}
