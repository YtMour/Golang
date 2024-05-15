package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义全局变量 票库存为10张
var ticket2 int = 10

// 定义一个锁  Mutex 锁头
var mutex sync.Mutex

func main() {
	// 单线程不存在问题，多线程资源争抢就出现了问题
	go saleTickets2("张三")
	go saleTickets2("李四")
	go saleTickets2("王五")
	go saleTickets2("赵六")

	time.Sleep(time.Second * 5)
}

// 售票函数
func saleTickets2(name string) {
	for {
		// 在拿到共享资源之前先上锁
		mutex.Lock()
		if ticket2 > 0 {
			time.Sleep(time.Millisecond * 1)
			fmt.Println(name, "剩余票的数量为：", ticket2)
			ticket2--
		} else {
			// 操作完毕后，解锁
			mutex.Unlock()
			fmt.Println("票已售完")
			break
		}
		// 操作完毕后，解锁
		mutex.Unlock()
	}
}
