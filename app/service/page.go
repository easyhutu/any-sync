package service

import "github.com/gin-gonic/gin"

func (svr *AnySyncSvr) Home(ctx *gin.Context) {
	ctx.HTML(200, "home.html", gin.H{
		"shareUrl": svr.config.ShareUrl,
	})
}

func (svr *AnySyncSvr) Mobile(ctx *gin.Context) {
	ctx.HTML(200, "mobile.html", gin.H{
		"shareUrl": svr.config.ShareUrl,
	})
}
