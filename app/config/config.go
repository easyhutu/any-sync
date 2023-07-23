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
		PingSecond:       time.Minute * 16, // 超过心跳间隔删除设备
		QrSize:           260,              // 分享二维码size
		MaxUploadFiles:   1025 << 20,       // 最大上传文件大小 Mib
		ShareFilesPrefix: "anySyncShare",   // 上传文件公共目录
	}
	config.BoundIp = device.GetBoundIp()
	config.ShareUrl = fmt.Sprintf("http://%s:%d/mobile", config.BoundIp, config.ListenPort)
	return config, nil
}
