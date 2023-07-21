package middleware

import (
	"any-sync/app/utils"
	"github.com/gin-gonic/gin"
)

const (
	UA    = "User-Agent"
	Buvid = "buvid"
)

func Dev() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		buvid, err := ctx.Cookie(Buvid)
		ua := ctx.GetHeader("User-Agent")
		if err != nil || buvid == "" {
			buvid = utils.ToMd5Str([]byte(ua))
			ctx.SetCookie(Buvid, buvid, 60*60*24*30, "", "", false, false)
			ctx.JSON(200, "set buvid")
			ctx.Abort()
			return
		}
		ctx.Set(Buvid, buvid)
		ctx.Set(UA, ua)
	}
}

func WithDevUA(ctx *gin.Context) string {
	return ctx.MustGet(UA).(string)
}

func WithDevBuvid(ctx *gin.Context) string {
	return ctx.MustGet(Buvid).(string)

}
