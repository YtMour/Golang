package main

import "fmt"

func main() {
	//调用函数
	var m int
	m = max(1, 2)
	fmt.Println(m)
}

// func 定义函数
func add1() {
	fmt.Println("hello world")
}
func max(num1, num2 int) int {
	var temp int
	if num2 > num1 {
		temp = num2
	} else {
		temp = num1
	}
	return temp
}
