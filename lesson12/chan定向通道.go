package main

import (
	"fmt"
	"time"
)

// 单向通道使用场景
func main() {
	//双向通道
	ch1 := make(chan int, 1) // 可读可写
	ch1 <- 100
	data := <-ch1
	fmt.Println(data)

	//	单向通道
	//ch2 := make(chan<- int, 1) //只能写,不能读
	//ch3 := make(<-chan int, 1) //只能读,不能写

	go writeOnly(ch1)
	go readOnly(ch1)
	time.Sleep(time.Second * 1)

}

// 作为函数的参数或者返回值之类的。
// 指定函数去写，不让他读取，防止通道滥用
func writeOnly(ch chan<- int) {
	// 函数的内部，处理一些写数据的操作
	ch <- 100
}

// 指定函数去读，不让他写，防止通道滥用
func readOnly(ch <-chan int) int {
	// 取出通道的值，做一些操作，不可写的。
	data := <-ch
	fmt.Println(data)
	return data
}
