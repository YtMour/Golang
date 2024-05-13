package main

import "fmt"

type AA interface {
	t1()
}
type BB interface {
	t2()
}

// 接口嵌套
type CC interface {
	AA // 导入AA 接口中的方法
	BB
	t3()
}

// 编写一个结构体实现接口CC
type d7 struct {
}

func (dog d7) t1() {
	fmt.Println("t1")

}
func (dog d7) t2() {
	fmt.Println("t2")

}
func (dog d7) t3() {
	fmt.Println("t3")

}

//编写一个结构体实现接口cc

func main() {
	var dog d7 = d7{}
	dog.t1()
	dog.t2()
	dog.t3()
	var a AA = dog
	a.t1()
}
