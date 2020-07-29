package controller

import (
	"github.com/gin-gonic/gin"
	"lastsummer/internal/service"
	"lastsummer/resp"
	"log"
	"os"
)

func ViewMyFile(ctx *gin.Context) {
	files, err := service.ViewMyFile(ctx)
	if err != nil {
		resp.Badmsg(ctx, "err in finding files")
	}
	ctx.JSON(200, files)
}

func Share(ctx *gin.Context) {
	filename := ctx.PostForm("filename")
	if !service.CheckPrivateFile(ctx, filename) {
		resp.Badmsg(ctx, "err in finding files,maybe it's not your file")
		return
	}
	link := service.GetShareLink(filename)
	ctx.String(200, link)
}

//
func ShareWithQRcode(ctx *gin.Context) {
	filename := ctx.PostForm("filename")
	if !service.CheckPrivateFile(ctx, filename) {
		resp.Badmsg(ctx, "err in finding files,maybe it's not your file")
		return
	}
	pic := service.GetShareQRcode(filename)
	ctx.File(pic)
	err := os.Remove(pic)
	if err != nil {
		log.Print(err)
	}
}

func CheckPrivate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//key := ctx.PostForm("key")
	}
}
