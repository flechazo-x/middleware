package work_queues

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq/static"
)

func GetQuery(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		static.WorkQueue, // 名称
		false,            // 持久
		false,            // 未使用时删除
		false,            // 独占
		false,            // 不等待
		nil,              // 参数
	)
}
