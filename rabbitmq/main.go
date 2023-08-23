package main

import (
	"rabbitmq/mq"
	"rabbitmq/pub_sub"
)

func main() {
	defer mq.Close()
	//hello_world.Producer()
	//hello_world.Consumer()
	//work_queues.Producer()
	//work_queues.Consumer()
	//pub_sub.Producer()
	pub_sub.Consumer()
	//rpc.Client()
	//rpc.Server()
}
