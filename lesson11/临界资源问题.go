package main

import (
	"fmt"
	"time"
)

// 定义全局变量 票库存为10张
var ticket int = 10

func main() {
	// 单线程不存在问题，多线程资源争抢就出现了问题
	go saleTickets("张三")
	go saleTickets("李四")
	go saleTickets("王五")
	go saleTickets("赵六")

	time.Sleep(time.Second * 5)
}

// 售票函数
func saleTickets(name string) {
	for {
		if ticket > 0 {
			time.Sleep(time.Millisecond * 1)
			fmt.Println(name, "剩余票的数量为：", ticket)
			ticket--
		} else {
			fmt.Println("票已售完")
			break
		}
	}
}
