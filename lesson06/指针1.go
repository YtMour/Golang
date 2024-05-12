package main

import "fmt"

// 指针
func main() {

	var a int = 10
	var b = &a
	fmt.Println("a变量的地址：", &a)
	fmt.Println("b存储的地址：", b)
	fmt.Println("b自己的地址：", &b)

	// &  *
	// b 指向的地址 ，  * 取出 b变量所指向的地址 中的值。
	fmt.Println("b变量指向的地址中的值：", *b)
	*b = 20
	fmt.Println(a) // a 变量20，通过指针变量b，操控了a的值。
	fmt.Println(*b)
}
