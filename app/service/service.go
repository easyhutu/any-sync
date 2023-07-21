package service

import (
	"any-sync/app/config"
	"any-sync/app/model/device"
	"bytes"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Home(ctx *gin.Context)
	Mobile(ctx *gin.Context)
	QrCreate(ctx *gin.Context)
	DevicePing(ctx *gin.Context)
	Sync(ctx *gin.Context)
	Upload(ctx *gin.Context)
	Download(ctx *gin.Context)
}

type AnySyncSvr struct {
	config   *config.Config
	devs     *device.Devices
	muFileHs map[string]map[int]*bytes.Buffer
}

func NewAnySyncSvr(config *config.Config) Service {
	svr := &AnySyncSvr{}
	svr.config = config
	svr.devs = device.NewDevices(config.PingSecond)
	svr.muFileHs = make(map[string]map[int]*bytes.Buffer)
	return svr
}
