package routing

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq/static"
)

func GetQuery(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // 名称
		false, // 持久
		false, // 未使用时删除
		true,  // 独占队列（当前声明队列的连接关闭后即被删除）
		false, // 不等待
		nil,   // 参数
	)
}

// CreateExchange 创建交换机
func CreateExchange(ch *amqp.Channel, name string) error {
	return ch.ExchangeDeclare(
		name,          // 交换机name
		static.Direct, // 交换机type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
}
