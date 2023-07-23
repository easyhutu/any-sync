package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

// Manager 所有 websocket 信息
type Manager struct {
	Group                   map[string]map[string]*Client
	groupCount, clientCount uint
	Lock                    sync.Mutex
	Register, UnRegister    chan *Client
	Message                 chan *MessageData
	GroupMessage            chan *GroupMessageData
	BroadCastMessage        chan *BroadCastMessageData
	MsgExecutor             MsgExecutor
}

// Client 单个 websocket 信息
type Client struct {
	Id, Group, Ua string
	Socket        *websocket.Conn
	Message       chan []byte
}

// MessageData messageData 单个发送数据信息
type MessageData struct {
	Id, Group string
	Message   []byte
}

// GroupMessageData groupMessageData 组广播数据信息
type GroupMessageData struct {
	Group   string
	Message []byte
}

// BroadCastMessageData 广播发送数据信息
type BroadCastMessageData struct {
	Message []byte
}
