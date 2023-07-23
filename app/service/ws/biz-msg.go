package ws

type MsgType string

type MsgExecutor func(client *Client, message *BizMessage)

const (
	PingMsg   = MsgType("ping")
	PingGroup = MsgType("pingGroup")
	SyncMsg   = MsgType("sync")
	IgnoreMsg = MsgType("ignore")
)

// BizMessage ws 接收到的业务 msg信息
type BizMessage struct {
	Type    MsgType     `json:"type,omitempty"`
	Content string      `json:"content,omitempty"`
	Val     interface{} `json:"val,omitempty"`
}

func WithMsgType(val string) MsgType {
	switch val {
	case string(PingMsg):
		return PingMsg
	case string(SyncMsg):
		return SyncMsg
	default:
		return IgnoreMsg
	}
}
