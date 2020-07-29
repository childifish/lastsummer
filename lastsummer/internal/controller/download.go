package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lastsummer/internal/service"
	"lastsummer/resp"
)

func Download(ctx *gin.Context) {
	filename := ctx.PostForm("filename")
	ok := service.CheckFileExist(filename)
	if !ok {
		resp.Badmsg(ctx, "unable to find file")
		return
	}
	ok = service.CheckPrivateFile(ctx, filename)
	if !ok {
		resp.Badmsg(ctx, "Maybe not your file")
		return
	}
	path := "./userfile/"
	filepath := path + filename
	fmt.Println(filepath)

	ctx.File(filepath)

}

func ShareDownLoad(ctx *gin.Context) {
	filename := ctx.Query("filename")
	ok := service.CheckFileExist(filename)
	if !ok {
		resp.Badmsg(ctx, "unable to find file")
		return
	}
	path := "./userfile/"
	filepath := path + filename
	ctx.File(filepath)
}
