package main

import (
	"fmt"
	"time"
)

// 定时器（time ： 当下，xxx之前before，xxxx之后 after）
func main() {
	// 创建一个定时器 NewTimer
	// Timer  C <-chan Time
	//timer := time.NewTimer(time.Second * 3)
	// 当前时间
	//fmt.Println(time.Now())
	// timer.C, 时间通道，这个通道中存放的值 2023-03-02 21:49:48.0453619 +0800 CST m=+3.013284501
	// 就是我们在定义定时器的时候，存放的时间，等待对应的时间。
	//timeChan := timer.C
	//fmt.Println(<-timeChan)

	// 定时器（提前关闭）
	// 会向Timer.C 放入一个时间。
	timer2 := time.NewTimer(time.Second * 5)
	go func() {
		<-timer2.C // 消费
		fmt.Println("end")
	}()
	timer2.Stop() // 手动停止定时器。
}
