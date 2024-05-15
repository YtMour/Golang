package main

import "fmt"

func main() {
	// 特殊的泛型类型，泛型的参数时多样的，但是实际类型定义就是int
	type AAA[T int | string] int

	var a AAA[int] = 123
	var b AAA[string] = 159
	//var c AAA[string] = "hello" // 底层类型是int
	fmt.Println(a)
	fmt.Println(b)
}
