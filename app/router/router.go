package router

import (
	"github.com/easyhutu/any-sync/app/config"
	"github.com/easyhutu/any-sync/app/service"
	"github.com/easyhutu/any-sync/app/utils/middleware"
	"github.com/gin-gonic/gin"
)

func RegRouter(engine *gin.Engine, cfg *config.Config) {
	svr := service.NewAnySyncSvr(cfg)
	svr.InitWebSocket()

	engine.GET("/", svr.Home)
	engine.GET("/mobile", svr.Mobile)
	engine.GET("/qr", svr.QrCreate)
	dev := engine.Group("/dev", middleware.Dev())
	{
		dev.GET("/ping", svr.DevicePing)
		dev.POST("/sync", svr.Sync)
		dev.POST("/upload", svr.Upload)
		dev.GET("/dl", svr.Download)
		dev.GET("/dl/anySyncShare/:fmd5/:filename", svr.Download)
		dev.GET("/ws/:channel", svr.SyncWebSocket) // websocket
	}

}
