package main

import "fmt"

func main() {
	// goroutine ： 和普通方法调用完全不同，它是并发执行的，快速交替。
	go hello()
	for i := 0; i < 100; i++ {
		fmt.Println("main - ", i)
	}
}
func hello() {
	for i := 0; i < 1000; i++ {
		fmt.Println("hello - ", i)
	}
}
