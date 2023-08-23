package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

var (
	mq   *amqp.Connection
	once sync.Once
)

// GetChannel 获取通道
func GetChannel() *amqp.Channel {
	once.Do(func() {
		conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
		if err != nil {
			panic(err)
		}
		mq = conn
	})
	ch, err := mq.Channel()
	if err != nil {
		log.Panicf("%s: %s", "无法打开通道", err)
		return nil
	}
	return ch
}

func Close() {
	mq.Close()
}
