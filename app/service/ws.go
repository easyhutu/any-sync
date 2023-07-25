package service

import (
	"github.com/easyhutu/any-sync/app/model/device"
	"github.com/easyhutu/any-sync/app/service/ws"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"time"
)

func (svr *AnySyncSvr) SyncWebSocket(ctx *gin.Context) {
	svr.wsManager.WsClient(ctx)
}

func (svr *AnySyncSvr) InitWebSocket() {
	go svr.wsManager.Start()
	go svr.wsManager.SendService()
	go svr.wsManager.SendService()
	go svr.wsManager.SendGroupService()
	go svr.wsManager.SendAllService()
}

func (svr *AnySyncSvr) revWsExecutor(client *ws.Client, message *ws.BizMessage) {
	log.Printf("executor [%s] message:%+v", client.Id, message)
	switch message.Type {
	case ws.PingMsg:
		svr.lock.Lock()
		svr.devs.Check(client.Id, client.Ua)
		svr.lock.Unlock()

		devs := device.NewDevices(svr.config.PingSecond)
		devs.Split(svr.devs, client.Id)
		smsg := &ws.BizMessage{
			Type: ws.SyncMsg,
			Val:  devs,
		}
		v, _ := json.Marshal(smsg)
		svr.wsManager.Send(client.Id, client.Group, v)

	case ws.PingGroup:
		svr.lock.Lock()
		svr.devs.Check(client.Id, client.Ua)
		svr.lock.Unlock()

		smsg := &ws.BizMessage{
			Type: ws.PingMsg,
		}
		v, _ := json.Marshal(smsg)
		svr.wsManager.SendGroup(client.Group, v)

	case ws.SyncMsg:
		params := &struct {
			ToMd5    string `json:"toMd5,omitempty"`
			SyncType string `json:"syncType,omitempty"`
			Content  string `json:"content,omitempty"`
			Desc     string `json:"desc,omitempty"`
		}{}
		err := json.Unmarshal([]byte(message.Content), params)
		if err != nil {
			log.Printf("json unmarshal err %v", err)
		}
		fromDev := svr.devs.WithDevice(client.Id)
		if fromDev == nil {
			return
		}
		si := &device.SyncInfo{
			From:     fromDev.Show,
			FromMd5:  fromDev.Md5,
			SyncTi:   time.Now(),
			Status:   device.SyncOK,
			Details:  []*device.SyncDetail{},
			SyncType: device.WithSyncType(params.SyncType),
		}
		si.Generate()

		sd := &device.SyncDetail{
			Content: params.Content,
			Desc:    params.Desc,
		}
		if si.SyncType == device.SyncTypeFile {
			sd = fromDev.FilterFile(params.Desc)
		}
		si.Details = append(si.Details, sd)
		log.Printf("sync info %+v", si)
		toDev := svr.devs.WithDevice(params.ToMd5)
		if toDev == nil {
			return
		}
		toDev.AddSync(*si)
		smsg := &ws.BizMessage{
			Type: ws.PingMsg,
		}
		v, _ := json.Marshal(smsg)

		// 通知目标设备发起ping
		svr.wsManager.Send(params.ToMd5, client.Group, v)
	}

	return
}
