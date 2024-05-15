package main

import (
	"fmt"
	"strconv"
	"time"
)

// 缓冲通道 chan，cap
func main() {

	// 非缓冲通道
	ch1 := make(chan int)
	fmt.Println(cap(ch1), len(ch1)) // 0 0
	//ch1 <- 100

	// 缓冲通道
	// 缓冲区通道，放入数据，不会产生死锁，它不需要等待另外的线程来拿，它可以放多个数据。
	// 如果缓冲区满了，还没有人取，也会产生死锁。
	ch2 := make(chan string, 5)
	fmt.Println(cap(ch2), len(ch2)) // 5 0
	ch2 <- "1"
	fmt.Println(cap(ch2), len(ch2)) // 5 1 , 可以通过len来判断缓冲通道中的数据数量
	ch2 <- "2"
	ch2 <- "3"
	fmt.Println(cap(ch2), len(ch2)) // 5 3
	ch2 <- "4"
	ch2 <- "5"
	fmt.Println(cap(ch2), len(ch2)) // 5 5
	data := <-ch2
	ch2 <- "6" // deadlock!
	fmt.Println(cap(ch2), len(ch2))
	fmt.Println(data)

	ch3 := make(chan string, 4)
	go test8(ch3)
	fmt.Println("--------------------------")
	for s := range ch3 {
		time.Sleep(time.Second)
		fmt.Println("main中读取的数据:", s)
	}
	fmt.Println("main-end")
}

func test8(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "test - " + strconv.Itoa(i)
		fmt.Println("子goroutine放入数据：", "test - "+strconv.Itoa(i))
	}
	close(ch)
}
