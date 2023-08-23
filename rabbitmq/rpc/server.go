package rpc

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/mq"
	"rabbitmq/utils"
	"strconv"
)

func Server() {
	//1.获取链接通道
	ch := mq.GetChannel()
	defer ch.Close()
	//2.声明队列
	q, err := GetQuery(ch)
	utils.FailOnError(err, "无法声明队列")

	err = ch.Qos(1, 0, false)
	utils.FailOnError(err, "未能设置 QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "无法消费消息")
	var forever chan struct{}
	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			utils.FailOnError(err, "Failed to convert body to integer")

			log.Printf(" [.] fib(%d)", n)

			err = ch.PublishWithContext(nil,
				"",        // 交换
				d.ReplyTo, // 路由键
				false,     // 强制
				false,     // 立即
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(fib(n))),
				})
			utils.FailOnError(err, "Failed to publish a message")
			_ = d.Ack(false)
		}
	}()
	<-forever
}

// 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

//go c mysql redis json/protobuf rabbitmq kafaka es gin pprof
