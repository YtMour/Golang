package main

import "fmt"

// 匿名结构体：没有名字的结构体
type Student struct {
	name string
	age  int
}

// 结构体中的匿名字段
type Teacher struct {
	string
	int
	// 如何打印这个匿名字段，默认使用数据类型当做字段名称
}

func main() {
	s1 := Student{"yt", 18}
	fmt.Println(s1)

	//	匿名结构体 直接可以在函数内部创建出来，创建后就需要赋值使用
	s2 := struct {
		name string
		age  int
	}{"iuiou", 90}
	fmt.Println(s2)
	//	匿名字段
	t1 := Teacher{"qqqq", 18}
	fmt.Println(t1)

}
