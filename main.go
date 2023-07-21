package main

import (
	"any-sync/app/config"
	"any-sync/app/router"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

var (

	//go:embed temps/*
	htmlFs embed.FS

	//go:embed static/*
	staticFs embed.FS
)

func main() {

	cfg, _ := config.NewConfig()
	engine := gin.Default()
	engine.Delims("{[{", "}]}")
	engine.SetHTMLTemplate(template.Must(template.New("").Delims("{[{", "}]}").ParseFS(htmlFs, "temps/*")))
	engine.Any("/static/*filepath", func(context *gin.Context) {
		staticServer := http.FileServer(http.FS(staticFs))
		staticServer.ServeHTTP(context.Writer, context.Request)
	})
	engine.MaxMultipartMemory = cfg.MaxUploadFiles
	println(engine.MaxMultipartMemory)
	router.RegRouter(engine, cfg)
	engine.Run(fmt.Sprintf(":%d", cfg.ListenPort))
}
