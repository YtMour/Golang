package main

import "fmt"

// 一个结构体可能包含一个字段，而这个字段又是一个结构体：结构体嵌套
type Preson struct {
	name    string
	age     int
	address Address
}
type Address struct {
	city, steate string
}

// 结构体是可以嵌套的 ：我们就可以定义很多复杂的对象，来进行拼接了，构成一个更大的对象
func main() {
	var p = Preson{
		name:    "asfs",
		age:     11,
		address: Address{"c", "s"},
	}
	fmt.Println(p)
}
