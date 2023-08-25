package routing

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/mq"
	"rabbitmq/static"
	"rabbitmq/utils"
	"time"
)

func Producer() {
	//1.获取链接通道
	ch := mq.GetChannel()
	defer ch.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exchangeName := "logs_direct" //交换机名称

	//2.创建交换机
	err := CreateExchange(ch, exchangeName)
	utils.FailOnError(err, "无法创建交换机")

	utils.ProducerRelease(func(input string) {
		//3.发布消息
		err = ch.PublishWithContext(ctx,
			exchangeName,        // 交换
			severityFrom(input), // 路由键
			false,               // 强制
			false,               // 立即
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(input),
			})
		if err != nil {
			log.Fatalf("%s: %s", "无法发布消息", err)
		} else {
			log.Printf(" %s 已发送 %s\n", static.WorkQueue, input)
		}
	})
}

func severityFrom(input string) string {
	if len(input) < 4 {
		fmt.Println("此条发布信息的等级是:info")
		return "info"
	}
	var ty string
	severity := input[:4]
	switch severity {
	case "info":
		ty = "info"
	case "warn":
		ty = "warn"
	default:
		ty = "error"
	}
	fmt.Println("此条发布信息的等级是:", ty)
	return ty
}
