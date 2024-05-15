package main

import "fmt"

// int8 衍生类型
type int8A int8
type int8B = int8

// ~ 表示可以匹配该类型的衍生类型
type NewInt interface {
	~int8
}

// ~
func main() {
	var a int8A = 10
	var b int8A = 10
	fmt.Println(GetMax(a, b))
}
func GetMax[T NewInt](a, b T) T {
	if a > b {
		return a
	}
	return b
}
