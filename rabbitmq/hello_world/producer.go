package hello_world

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/mq"
	"rabbitmq/static"
	"rabbitmq/utils"
	"time"
)

// Producer 生产者
func Producer() {
	//1.获取链接通道
	ch := mq.GetChannel()
	defer ch.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//2.声明持久化队列
	q, err := GetQuery(ch)
	utils.FailOnError(err, "无法声明队列")
	utils.ProducerRelease(func(input string) {
		//3.发布消息
		err = ch.PublishWithContext(ctx,
			"",     // 交换
			q.Name, // 路由键
			false,  // 强制
			false,  // 立即
			amqp.Publishing{
				DeliveryMode: amqp.Persistent, // 持久（交付模式：瞬态/持久）
				ContentType:  "text/plain",
				Body:         []byte(input),
			})
		if err != nil {
			log.Fatalf("%s: %s", "无法发布消息", err)
		} else {
			log.Printf(" %s 已发送 %s\n", static.HelloWorld, input)
		}
	})
}
