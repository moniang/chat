# 聊天程序
目前实现的功能为：

- 登录注册
- 群聊
- 修改昵称
- 修改字体颜色
- 上线提示
- VIP过场动画

### 用到的第三方包有
- github.com/nsqio/go-nsq
- github.com/gorilla/websocket
- github.com/jinzhu/gorm
- github.com/moniang/validate

## 使用说明
首次使用，请导入SQL文件
先启动NSQ
```
nsqlookupd
nsqd --lookupd-tcp-address=127.0.0.1:4160
```
然后启动程序
``` 
go build
chat
```
界面：http://127.0.0.1:8080/
界面在Vue分支

演示地址：https://c.imoniang.com

### 感谢
感谢邪少大佬远程教会我如何使用Git，还帮我搭建了服务器环境

感谢，二萌，静看花开花落，Zero，唐吉诃德，孤独的观测者，顾飞惜，我是一名剑客，硅钠砝，人间四月等对萌新墨娘的一切帮助！

感谢Go/Golang语言开发群（780349811）群友的帮助

### 谢谢你们
