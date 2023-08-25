package rpc

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math/rand"
	"rabbitmq/mq"
	"rabbitmq/utils"
	"strconv"
	"time"
)

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func fibonacciRPC(input string) (res int, err error) {
	ch := mq.GetChannel()
	defer ch.Close()

	q, err := GetQuery(ch, "", true)
	utils.FailOnError(err, "无法声明队列")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "消费者注册失败")

	corrId := randomString(32)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(input),
		})
	utils.FailOnError(err, "发布消息失败")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			utils.FailOnError(err, "无法将正文转换为整数")
			break
		}
	}

	return
}

func Client() {
	rand.Seed(time.Now().UTC().UnixNano())
	utils.ProducerRelease(
		func(s string) {
			res, err := fibonacciRPC(s)
			utils.FailOnError(err, "无法处理 RPC 请求")
			log.Printf(" [.] 收到服务器返回结果 %d", res)
		},
	)

}
