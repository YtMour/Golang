package main

import "fmt"

// defer 传参的调用时机
func main() {
	n := 10
	fmt.Println("main中n=", n)
	// 在defer的时候，函数其实已经被调用了，但是没有执行。参数是已经传递进去的了
	defer ff(n)
	n += 15
	fmt.Println("mian end n=", n)
}
func ff(n int) {
	fmt.Println("函数中n=", n)
}
