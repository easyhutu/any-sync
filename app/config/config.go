package config

import (
	"any-sync/app/utils/device"
	"fmt"
	"time"
)

type Config struct {
	BoundIp          string
	ListenPort       int
	PingSecond       time.Duration
	ShareUrl         string
	QrSize           int
	MaxUploadFiles   int64
	ShareFilesPrefix string
}

func NewConfig() (config *Config, err error) {
	config = &Config{
		ListenPort:       8080,
		PingSecond:       time.Second * 10,
		QrSize:           260,
		MaxUploadFiles:   260 << 20, // 260 Mib
		ShareFilesPrefix: "anySyncShare",
	}
	config.BoundIp = device.GetBoundIp()
	config.ShareUrl = fmt.Sprintf("http://%s:%d/mobile", config.BoundIp, config.ListenPort)
	return config, nil
}
