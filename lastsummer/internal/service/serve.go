package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"lastsummer/internal/middleware"
	"lastsummer/internal/model"
	"lastsummer/resp"
	"os"
)

func ViewMyFile(ctx *gin.Context) ([]model.File, error) {
	claim, ok := ctx.Get("claim")
	if !ok {
		return nil, errors.New("中间件读取错误")
	}

	var u model.User
	u.Username = claim.(*middleware.Claims).AppKey

	var f model.File
	files, err := f.FindMyFile(u.GetUid())
	if err != nil {
		return nil, err
	}
	return files, nil
}

func CheckFileExist(filename string) bool {
	filepath := "./userfile/" + filename
	fmt.Println(filepath)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckPrivateFile(ctx *gin.Context, filename string) bool {
	files, err := ViewMyFile(ctx)
	if err != nil {
		resp.Badmsg(ctx, "err in finding files")
	}
	filepath := "./userfile/" + filename

	for _, v := range files {
		if v.FilePath == filepath {
			return true
		}
	}
	return false
}
