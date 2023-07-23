# any sync

局域网同步工具，不需要登录QQ、微信即可在局域网内跨平台同步数据

> PC端

![screen](/doc/pc.png)

> 移动端

![screen](/doc/phone.png)


# 使用方法

### [下载对应可执行文件](https://github.com/easyhutu/any-sync/releases)

### DEV

> 项目依赖 go 1.20

项目没有外漏配置参数，如需修改配置信息请查看 [/app/config/config.go](/app/config/config.go)


```shell
go mod tidy
go build -o anysync && anysync
```

***
#### *注意：项目未做任何加密和鉴权处理，使用中请注意隐私保护*