package lib

import (
	"encoding/json"
	"math/rand"
	"time"
)

type returnJson struct {
	Code int         `json:"code"` // 返回状态码
	Msg  string      `json:"msg"`  // 返回消息
	Data interface{} `json:"data"` // 附带数据
}

// 获取随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 返回请求Json
func MakeReturnJson(code int, message string, data interface{}) []byte {
	returnJson := &returnJson{}
	returnJson.Code = code
	returnJson.Msg = message
	returnJson.Data = data
	jsonStu, err := json.Marshal(returnJson)
	if err != nil {
		return []byte("")
	} else {
		return jsonStu
	}
}
