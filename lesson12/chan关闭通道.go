package main

import (
	"fmt"
	"time"
)

// 关闭通道
// 告诉接收方，我不会再有其他数据发送到chan了。
func main() {

	// 在main线程中定义的通道
	ch1 := make(chan int)
	go test7(ch1)
	// 读取chan中的数据
	//for {
	//	time.Sleep(time.Second)
	//	// ok 判断chan的状态是否是关闭，如果是关闭，不会再取值了。
	//	// ok, 如果是true，就代表我们还在读数据
	//	// ok, 如果是fasle，就说明该通道已关闭
	//	data, ok := <-ch1
	//	if !ok {
	//		fmt.Println("读取完毕", ok)
	//		break
	//	}
	//	fmt.Println("ch1 data:", data)
	//}

	// 读取chan中的数据, for 一个个取，并且会自动判断chan是否close 迭代器
	for data := range ch1 {
		time.Sleep(time.Second)
		fmt.Println(data)
	}
	fmt.Println("end")
}

// 通道可以参数传递
func test7(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// 关闭通道，告诉接收方，不会在往ch中放入数据
	close(ch)
}
