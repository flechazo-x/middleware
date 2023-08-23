package rpc

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq/static"
)

func GetQuery(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"rpc_queue", // 名称
		false,       // 持久
		false,       // 未使用时删除
		true,        // 独占队列（当前声明队列的连接关闭后即被删除）
		false,       // 不等待
		nil,         // 参数
	)
}

func CreateExchange(ch *amqp.Channel, name string) error {
	return ch.ExchangeDeclare(
		name,          // name
		static.Fanout, // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments

	)
}
