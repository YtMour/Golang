package main

import "fmt"

// 泛型的约束提取定义
type MyInt interface {
	int | int8 | int16 | int32 | int64 // 作用域泛型的，而不是一个接口方法
}

// 自定义泛型
func main() {
	var a = 15
	var b = 20
	GetMaxNum(a, b)
	fmt.Println(GetMaxNum(a, b))
}
func GetMaxNum[T MyInt](a, b T) T {
	if a > b {
		return a
	}
	return b
}
