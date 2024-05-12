package main

import "fmt"

func main() {
	f12()
	f2 := f12
	f2()
	//f12和f2 本质指向了同一个内存空闲
	//匿名函数   在函数体后增加（） 调用这个函数，匿名函数只能用一次
	func() {
		fmt.Println("我是一个匿名函数")
	}()

	//将匿名函数进行赋值  就可以实现多次调用
	f3 := func() {
		fmt.Println("我是一个匿名函数  2")
	}
	f3()

	//	匿名函数是否可以添加参数和返回值
	func(a, b int) {
		fmt.Println("a,b=", a, b)
	}(1, 2)

	//	匿名函数返回值定义给变量
	r1 := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(r1)

	// 由于Go语言中的函数是一个特殊的变量，支持匿名操作
	// Go语言支持函数式编程
	// - 将匿名函数作为另外一个函数的参数，回调函数
	// - 将匿名函数作为另外一个函数的返回值，可以形成闭包结构
}

func f12() {
	fmt.Println("我是f12函数")
}
