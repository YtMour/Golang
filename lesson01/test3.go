package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println(a, b)
	c := 60
	return c
}

// 返回多个返回值  匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a=", a, "b=", b)
	return 77, 88
}

// 返回多个返回值  有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---foo3---")
	fmt.Println("a=", a, "b=", b)
	//给有名称的返回值变量赋值
	r1 = 123
	r2 = 456
	return
}
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---foo4---")
	fmt.Println("a=", a, "b=", b)
	//给有名称的返回值变量赋值
	r1 = 1
	r2 = 2
	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println(c)

	var ret1, ret2 = foo2("hhh", 999)
	fmt.Println("ret1=", ret1, "ret2=", ret2)
	ret1, ret2 = foo3("foo3", 159)
	fmt.Println("ret1=", ret1, "ret2=", ret2)
	ret1, ret2 = foo4("foo4", 0)
	fmt.Println("ret1=", ret1, "ret2=", ret2)

}
