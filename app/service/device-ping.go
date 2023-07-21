package service

import (
	"any-sync/app/model/device"
	"any-sync/app/utils/middleware"
	"github.com/gin-gonic/gin"
)

func (svr *AnySyncSvr) DevicePing(ctx *gin.Context) {
	ua := middleware.WithDevUA(ctx)
	buvid := middleware.WithDevBuvid(ctx)
	if ua != "" {
		svr.devs.Check(buvid, ua)
	}
	devs := device.NewDevices(svr.config.PingSecond)
	devs.Split(svr.devs, buvid)

	ctx.JSON(200, devs)

}
