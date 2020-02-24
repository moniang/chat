package main

import (
	"github.com/moniang/chat/service"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"time"
)

var producer *nsq.Producer

// 消费者
type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	service.HandleMessage(msg)
	return nil
}

// 初始化生产者
func InitProducer(ipAddr string) (errorNo int, err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(ipAddr, config)
	if err != nil {
		return 1, err
	}

	producer.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags), nsq.LogLevelInfo)
	err = producer.Ping()
	if err != nil {
		return 2, err
	}
	return 0, nil
}

// 初始化消费者
func InitConsumer(topic string, channel string, address string) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = time.Second // 设置重连时间
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, 0)        // 屏蔽系统日志
	c.AddHandler(&ConsumerT{}) // 增加消费者接口

	if err := c.ConnectToNSQLookupd(address); err != nil { // 建立NSQLookupd连接
		panic(err)
	}
}

// 发送消息至NSQ
func SendMessage(topic string, Msg []byte) error {
	return producer.Publish(topic, Msg)
}
