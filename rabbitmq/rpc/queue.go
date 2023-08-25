package rpc

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func GetQuery(ch *amqp.Channel, name string, exclusive bool) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,      // 名称
		false,     // 持久
		false,     // 未使用时删除
		exclusive, // 独占队列（当前声明队列的连接关闭后即被删除）
		false,     // 不等待
		nil,       // 参数
	)
}
