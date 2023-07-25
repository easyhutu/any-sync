package main

import (
	"any-sync/app/config"
	"any-sync/app/router"
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"time"
)

type SrvStatus string

const (
	SrvRunning = SrvStatus("running")
	SrvStop    = SrvStatus("stop")
)

var (

	//go:embed temps/*
	htmlFs embed.FS

	//go:embed static/*
	staticFs embed.FS
)

type AnySyncServer struct {
	srv    *http.Server
	Status SrvStatus
	cfg    *config.Config
}

func NewAnySyncServer() *AnySyncServer {
	gin.SetMode(gin.ReleaseMode)
	cfg, _ := config.NewConfig()
	engine := gin.Default()

	engine.SetHTMLTemplate(template.Must(template.New("").Delims("{[{", "}]}").ParseFS(htmlFs, "temps/*")))
	engine.Any("/static/*filepath", func(context *gin.Context) {
		staticServer := http.FileServer(http.FS(staticFs))
		staticServer.ServeHTTP(context.Writer, context.Request)
	})
	engine.MaxMultipartMemory = cfg.MaxUploadFiles
	router.RegRouter(engine, cfg)
	return &AnySyncServer{
		srv: &http.Server{
			Handler: engine,
			Addr:    fmt.Sprintf(":%d", cfg.ListenPort),
		},
		Status: SrvStop,
		cfg:    cfg,
	}
}

func (s *AnySyncServer) Run() bool {
	if s.Status == SrvRunning {
		log.Println("server is running...")
		return true
	}

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Status = SrvStop
			log.Fatalf("listen: %s\n", err)
		}

	}()
	log.Printf("server start http://%s:%d", s.cfg.BoundIp, s.cfg.ListenPort)
	s.Status = SrvRunning
	return true
}

func (s *AnySyncServer) Stop() bool {
	if s.Status == SrvStop {
		log.Println("server is stop...")
		return true
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		s.Status = SrvRunning
		log.Fatal("Server forced to shutdown: ", err)
	}
	s.Status = SrvStop
	log.Println("Server exiting")
	return true
}
