package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 最小间隔 2秒，随机区间 3秒，即最大值为 2+3=5秒
	minDuration := int(2 * time.Second)
	randDuration := int(3 * time.Second)

	duration := time.Duration(rand.Intn(randDuration) + minDuration)

	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			fmt.Println("此次间隔：", duration)
			duration = time.Duration(rand.Intn(randDuration) + minDuration)
			timer.Reset(duration)
		default:
			// 避免 select 阻塞当前线程，具体查询 select 语法
		}
	}
}
