package routing

import (
	"log"
	"os"
	"rabbitmq/mq"
	"rabbitmq/utils"
)

func Consumer() {
	//1.获取链接通道
	ch := mq.GetChannel()
	defer ch.Close()

	exchangeName := "logs_direct" //交换机名称
	//2.创建交换机
	err := CreateExchange(ch, exchangeName)
	utils.FailOnError(err, "无法创建交换机")

	//3.声明队列
	q, err := GetQuery(ch)
	utils.FailOnError(err, "无法声明队列")

	//4.绑定队列到交换机
	for _, s := range os.Args[1:] {
		log.Printf("绑定队列:%s到交换机,\n感兴趣的内容是: %s", q.Name, s)
		err = ch.QueueBind(q.Name, s, exchangeName, false, nil)
		utils.FailOnError(err, "队列绑定失败")
	}

	utils.FailOnError(err, "无法绑定队列到交换机")

	msgs, err := ch.Consume(
		q.Name, // 队列
		"",     // 消费者
		true,   // 自动确认
		false,  // 独占
		false,  // 非本地
		false,  // 无等待
		nil,    // args
	)
	utils.FailOnError(err, "无法消费消息")
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("收到消息: %s", d.Body)
		}
	}()
	<-forever
}
