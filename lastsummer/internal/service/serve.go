package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"lastsummer/internal/middleware"
	"lastsummer/internal/model"
	"lastsummer/resp"
	"os"
	"strconv"
	"time"
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

//创建url
func GetShareLink(filename string) string {
	//url := "http://118.31.20.31:8080/sharelink?filename=" + filename
	url := "localhost:8012/sharelink?filename=" + filename
	return url
}

//这里用时间戳避免重复
//可以将时间戳写入表，用来判断链接过期时间
func GetShareQRcode(filename string) string {
	//url := "http://118.31.20.31:8080/sharelink?filename=" + filename
	url := "localhost:8012/sharelink?filename=" + filename
	pic := "./" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"
	qrcode.WriteFile(url, qrcode.Medium, 256, pic)
	return pic
}
