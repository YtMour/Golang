package main

import (
	"fmt"
	"time"
)

// select
func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 200
	}()

	// 读取chan数据，无论谁先放入，我们就用谁，抛弃其他的.
	// select 和 swtich 只是在通道中使用，case表达式需要是一个通道结果
	select {
	case num1 := <-ch1:
		fmt.Println(num1)
	case num2 := <-ch2:
		fmt.Println(num2)
		//default:
		// fmt.Println("default")
	}

}
