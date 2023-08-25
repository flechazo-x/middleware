package rpc

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/mq"
	"rabbitmq/utils"
	"strconv"
	"time"
)

func square(n int) int {
	return n * n
}

func Server() {
	ch := mq.GetChannel()
	defer ch.Close()

	q, err := GetQuery(ch, "rpc_queue", false)
	utils.FailOnError(err, "无法声明队列")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	utils.FailOnError(err, "设置QoS失败")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "消费者注册失败")

	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			utils.FailOnError(err, "无法将正文转换为整数")

			log.Printf(" [.] 计算请求参数(%d)的平方", n)
			response := square(n)

			err = ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			utils.FailOnError(err, "发布消息失败")

			_ = d.Ack(false)
		}
	}()
	log.Printf(" [*] 等待 RPC 请求")
	<-forever
}
