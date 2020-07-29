package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lastsummer/internal/middleware"
	"lastsummer/internal/model"
	"log"
	"mime/multipart"
	"strconv"
	"time"
)

func SaveFile(file *multipart.FileHeader, ctx *gin.Context) error {
	path := "./userfile/"
	filename := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err := ctx.SaveUploadedFile(file, path+filename)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(path + filename)
	var f model.File

	claim, ok := ctx.Get("claim")
	if !ok {
		return errors.New("中间件读取错误")
	}

	var u model.User
	u.Username = claim.(*middleware.Claims).AppKey

	err2 := f.WriteFileOrigin(u.GetUid(), path+filename)
	if err2 != nil {
		log.Println(err2)
		return err2
	}

	return nil
}
