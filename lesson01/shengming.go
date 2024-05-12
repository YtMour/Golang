package main

import "fmt"

func main() {
	//指定变量类型 声明后不赋值 使用默认值 int的默认值为0
	var i int
	fmt.Println(i)

	//根据值 自行判断变量类型（类型推导）
	var num = 10
	fmt.Println(num)

	//一次性声明多个变量
	var n1, n2, n3, name = 1, 2.2, 3, "Yt"
	fmt.Println(n1, n2, n3, name)

	var (
		a1  = 100
		a2  = 200
		att = "tryyui"
	)
	fmt.Println(a1, a2, att)
}
