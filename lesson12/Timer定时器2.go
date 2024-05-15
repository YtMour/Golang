package main

import (
	"fmt"
	"time"
)

func main() {
	// chan 放入当前时间之后的某个时间
	//timerChan := time.After(time.Second * 3)
	//fmt.Println(time.Now())
	//chanTime := <-timerChan
	//fmt.Println(chanTime)

	// 在3s以后执行这个函数
	time.AfterFunc(time.Second*3, mail)

	time.Sleep(5 * time.Second)
}
func mail() {
	// 发邮件
	fmt.Println("发邮件")
}
