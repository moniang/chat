package service

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// 客户端
type Client struct {
	Token     string          // 用户Token，用来请求信息
	Id        int             // 用户ID
	Conn      *websocket.Conn // 用户webSocket连接
	Name      string          // 用户昵称
	FontColor string          // 字体颜色
	UpTime    int64           //上线时间
}

// 消息体结构
type Message struct {
	ID        int    `json:"id"`         // 发送者ID
	Nick      string `json:"nick"`       // 发送者昵称
	Message   string `json:"message"`    // 消息内容
	SendTime  int64  `json:"send_time"`  // 发送时间(时间戳)
	FontColor string `json:"font_color"` // 字体颜色
}

// 跨域配置
var upgrader websocket.Upgrader

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool { // 允许跨域请求
			return true
		},
	}
}
