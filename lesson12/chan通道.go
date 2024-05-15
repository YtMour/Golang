package main

import (
	"fmt"
	"time"
)

// 定义通道 chan
// 这个 goroutine 希望告诉 main 线程，我还没结束。（通信）
func main() {
	// 定一个bool的通道
	var ch chan bool
	ch = make(chan bool)

	// 在一个goroutine中去往通道中放入数据
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("goroutine-", i)
		}
		time.Sleep(time.Second * 3)
		ch <- true
	}()

	// 另一个goroutine可以从通道中取出数据。（线程之间的通信）
	// 阻塞等待ch拿到值。有另外一个goroutine往里放值。
	data := <-ch
	fmt.Println("ch data:", data)
}
