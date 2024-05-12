package main

import "fmt"

// 就近原则   先启用全局变量
// Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑。
// string  int   float64  浮点数（小数）
var a float64 = 3.14

func main() {
	var a float64 = 2.18
	fmt.Println(a)
}
