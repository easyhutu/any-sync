package config

import (
	"fmt"
	"github.com/easyhutu/any-sync/app/utils/device"
	"time"
)

type Config struct {
	BoundIp          string
	ListenPort       int
	PingSecond       time.Duration // 超过心跳间隔删除设备
	ShareUrl         string
	QrSize           int    // 分享二维码size
	MaxUploadFiles   int64  // 最大上传文件大小 Mib
	ShareFilesPrefix string // 上传文件公共目录
}

func NewConfig() (config *Config, err error) {
	config = &Config{
		ListenPort:       8080,
		PingSecond:       time.Minute * 16,
		QrSize:           260,
		MaxUploadFiles:   1025 << 20,
		ShareFilesPrefix: "anySyncShare",
	}
	config.BoundIp = device.GetBoundIp()
	config.ShareUrl = fmt.Sprintf("http://%s:%d/mobile", config.BoundIp, config.ListenPort)
	return config, nil
}
