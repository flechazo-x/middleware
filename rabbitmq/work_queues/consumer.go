package work_queues

import (
	"bytes"
	"log"
	"rabbitmq/mq"
	"rabbitmq/utils"
	"time"
)

func Consumer() {
	//1.获取链接通道
	ch := mq.GetChannel()
	defer ch.Close()
	//2.声明队列
	q, err := GetQuery(ch)
	utils.FailOnError(err, "无法声明队列")

	err = ch.Qos(
		1,     // 预取计数
		0,     // 预取大小
		false, // 全局
	)
	utils.FailOnError(err, "未能设置 QoS")

	msgs, err := ch.Consume(
		q.Name, // 队列
		"",     // 消费者
		false,  // 自动确认
		false,  // 独占
		false,  // 非本地
		false,  // 无等待
		nil,    // args
	)
	utils.FailOnError(err, "无法消费消息")

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			_ = d.Ack(false)
			log.Printf("Done")
		}
	}()
	<-forever
}
