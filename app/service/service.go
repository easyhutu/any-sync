package service

import (
	"bytes"
	"github.com/easyhutu/any-sync/app/config"
	"github.com/easyhutu/any-sync/app/model/device"
	"github.com/easyhutu/any-sync/app/service/ws"
	"github.com/gin-gonic/gin"
	"sync"
)

type Service interface {
	Home(ctx *gin.Context)
	Mobile(ctx *gin.Context)
	QrCreate(ctx *gin.Context)
	DevicePing(ctx *gin.Context)
	Sync(ctx *gin.Context)
	Upload(ctx *gin.Context)
	Download(ctx *gin.Context)
	SyncWebSocket(ctx *gin.Context)
	InitWebSocket()
}

type AnySyncSvr struct {
	config    *config.Config
	devs      *device.Devices
	muFileHs  map[string]map[int]*bytes.Buffer
	wsManager *ws.Manager
	lock      sync.Mutex
}

func NewAnySyncSvr(config *config.Config) Service {
	svr := &AnySyncSvr{}
	svr.config = config
	svr.devs = device.NewDevices(config.PingSecond)
	svr.muFileHs = make(map[string]map[int]*bytes.Buffer)

	svr.wsManager = ws.NewWebsocketManager(svr.revWsExecutor)
	return svr
}
