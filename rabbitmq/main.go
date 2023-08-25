package main

import (
	"rabbitmq/mq"
	"rabbitmq/rpc"
)

func main() {
	defer mq.Close()
	//hello_world.Producer()
	//hello_world.Consumer()
	//work_queues.Producer()
	//work_queues.Consumer()
	//pub_sub.Producer()
	//pub_sub.Consumer()
	//routing.Producer()
	//routing.Consumer()
	//rpc.Client()
	rpc.Server()
	//topic.Producer()
	//topic.Consumer()
}
