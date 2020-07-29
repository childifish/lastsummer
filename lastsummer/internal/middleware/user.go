package middleware

import (
	"github.com/gin-gonic/gin"
	"lastsummer/resp"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			resp.Badmsg(ctx, "空token")
			ctx.Abort()
			return
		} else {
			claims, err := ParseToken(auth)
			if err != nil {
				resp.Badmsg(ctx, "有问题的token")
				ctx.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				resp.Badmsg(ctx, "过期token，请重新登陆")
				ctx.Abort()
				return
			}
			ctx.Set("claim", claims)
		}
		ctx.Next()
	}
}
