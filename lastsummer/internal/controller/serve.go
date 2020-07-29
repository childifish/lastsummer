package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"lastsummer/internal/service"
	"lastsummer/resp"
	"log"
	"os"
	"strconv"
	"time"
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
	link := GetShareLink(filename)
	ctx.String(200, link)
}

func GetShareLink(filename string) string {
	url := "http://118.31.20.31:8080/sharelink?filename=" + filename
	//url := "localhost:8080/sharelink?filename=" + filename
	return url
}

func GetShareQRcode(filename string) string {
	url := "http://118.31.20.31:8080/sharelink?filename=" + filename
	//url := "localhost:8080/sharelink?filename=" + filename
	pic := "./" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"
	qrcode.WriteFile(url, qrcode.Medium, 256, pic)
	return pic
}

func ShareWithQRcode(ctx *gin.Context) {
	filename := ctx.PostForm("filename")
	if !service.CheckPrivateFile(ctx, filename) {
		resp.Badmsg(ctx, "err in finding files,maybe it's not your file")
		return
	}
	pic := GetShareQRcode(filename)
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
