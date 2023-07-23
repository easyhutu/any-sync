package service

import (
	"any-sync/app/model/device"
	"any-sync/app/utils/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (svr *AnySyncSvr) Sync(ctx *gin.Context) {
	params := &struct {
		ToMd5    string `form:"toMd5"`
		SyncType string `form:"syncType"`
		Content  string `form:"content"`
		Desc     string `form:"desc"`
	}{}
	_ = ctx.ShouldBind(params)
	if params.ToMd5 == "" || params.SyncType == "" {
		ctx.JSON(400, "param err")
		return
	}
	fromBuvid := middleware.WithDevBuvid(ctx)
	fromDev := svr.devs.WithDevice(fromBuvid)
	println(fmt.Sprintf("hf:%+v, params: %+v", fromDev.HasFiles, params))

	si := &device.SyncInfo{
		From:     fromDev.Show,
		FromMd5:  fromDev.Md5,
		SyncTi:   time.Now(),
		Status:   device.SyncOK,
		Details:  []*device.SyncDetail{},
		SyncType: device.WithSyncType(params.SyncType),
	}
	si.Generate()

	sd := &device.SyncDetail{
		Content: params.Content,
		Desc:    params.Desc,
	}
	if si.SyncType == device.SyncTypeFile {
		sd = fromDev.FilterFile(params.Desc)
	}
	fmt.Printf("sync detail: %+v", sd)
	si.Details = append(si.Details, sd)
	svr.devs.WithDeviceMd5(params.ToMd5).AddSync(*si)
	ctx.JSON(200, "success")
}
