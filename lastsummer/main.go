package main

import (
	"github.com/gin-gonic/gin"
	"lastsummer/internal/controller"
	"lastsummer/internal/middleware"
	"lastsummer/internal/model"
)

func main() {
	model.InitDB()
	r := gin.Default()

	user := r.Group("/user")
	{
		//注册
		user.POST("/register", controller.Register)
		//登录
		user.POST("/login", controller.Login)
	}
	pan := r.Group("/pan", middleware.JWT())
	{
		//获取全部属于自己的文件的位置
		pan.GET("/myfile", controller.ViewMyFile)
		//生成分享链接
		pan.POST("/share", controller.Share)
		//生成二维码
		pan.POST("/shareqr", controller.ShareWithQRcode)
		//上传
		pan.POST("/upload", controller.UploadFile)
		//上传多文件
		pan.POST("/uploadmuti", controller.UploadMutiFile)
		//下载（Basic）
		pan.GET("/download", controller.Download)
	}
	//通过分享链接下载
	//secdownload := r.Group("/private",controller.)
	r.GET("/sharelink", controller.ShareDownLoad)

	r.Run(":8012")
}
