package controller

import (
	"github.com/gin-gonic/gin"
	"lastsummer/internal/service"
	"lastsummer/resp"
)

//最最基础的上传
func UploadFile(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	if err != nil {
		resp.Badmsg(ctx, "上传过程中出现问题")
		return
	}

	err2 := service.SaveFile(file, ctx)

	if err2 != nil {
		resp.Badmsg(ctx, "保存中出现问题")
		return
	}

	resp.Goodmsg(ctx, "上传成功")
}
