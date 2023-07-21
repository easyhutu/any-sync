package service

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func (svr *AnySyncSvr) QrCreate(ctx *gin.Context) {
	params := &struct {
		Txt string `form:"txt"`
	}{}
	_ = ctx.ShouldBind(params)
	qrstr := svr.config.ShareUrl
	var png []byte

	if params.Txt != "" {

		b64, _ := base64.StdEncoding.DecodeString(params.Txt)
		qrstr = string(b64)
	}
	println(qrstr)
	png, err := qrcode.Encode(qrstr, qrcode.Medium, svr.config.QrSize)
	if err != nil {
		println("create qr err", err.Error())
	}
	ctx.Writer.WriteHeader(200)
	ctx.Header("Content-Type", "image/png")
	ctx.Header("Accept-Length", fmt.Sprintf("%d", len(png)))
	ctx.Writer.Write(png)

}
