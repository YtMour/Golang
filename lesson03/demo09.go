package main

import "fmt"

// defer 作用：处理一些善后的问题 比如错误，文件，网络流关闭等
// 延迟函数   函数可以添加多个defer语句
// **当函数执行到最后时，这些defer语句会按照  逆序执行
func main() {
	defer f("yt")

	f("1")
	fmt.Println("2")
	defer f("3")
	fmt.Println("4")
}

func f(s string) {
	fmt.Println(s)
}
