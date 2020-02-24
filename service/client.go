package service

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"net/http"
	"sync"
	"time"
)

// 用户列表
var SocketList = &sync.Map{}

// 新建客户端连接
func NewSocketClient(token string, w http.ResponseWriter, r *http.Request) (client *Client) {
	conn, err := upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		return
	}

	client = &Client{
		Conn:   conn,
		Token:  token,
		UpTime: time.Now().Unix(),
	}
	return client
}

// 处理消息
func HandleMessage(msg *nsq.Message) {
	var m Message
	err := json.Unmarshal(msg.Body, &m)
	if err != nil {
		return
	}

	SocketList.Range(func(k, v interface{}) bool {
		client := v.(Client)
		if client.Id != m.ID {
			client.Conn.WriteJSON(m) // 自己的消息不发给自己
		}
		return true
	})

}
