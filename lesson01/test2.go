package main

import "fmt"

// const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota 每行的iota都会累加 1  第一行的iota的默认值是0
	//iota 只能配合 const 来使用
	BEIJNG = 10 * iota
	SHANGHAI
	SHENGZHEN
)

func main() {
	//常量(只读)
	const length int = 10

	fmt.Println(length)

	//length = 100  常量不允许修改

	fmt.Println(BEIJNG, SHANGHAI, SHENGZHEN)
}
