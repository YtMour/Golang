package main

import "fmt"

// 回调函数
func main() {
	r1 := add(1, 2)
	fmt.Println(r1)

	r2 := oper(1, 2, add)
	fmt.Println(r2)
	r3 := oper(1, 2, sub)
	fmt.Println(r3)

	//匿名函数
	fun1 := func(a, b int) int {
		return a * b
	}
	r4 := oper(4, 2, fun1) //调用匿名函数
	fmt.Println(r4)
	//	直接传递匿名函数
	r5 := oper(4, 2, func(a int, b int) int {
		if b == 0 {
			fmt.Println("不能为0")
			return 0
		}
		return a / b
	})
	fmt.Println(r5)
}

// 运算
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

// 高阶函数
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	r := fun(a, b)
	return r
}
