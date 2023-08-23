package utils

import (
	"fmt"
	"log"
	"time"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// ProducerRelease 生产者发布信息
func ProducerRelease(fn func(string)) {
	for {
		var input string
		fmt.Println("请输入发布内容：")
		fmt.Scanln(&input)
		fn(input)
		time.Sleep(time.Millisecond)
	}
}
