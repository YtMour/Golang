package main

import "fmt"

// 函数是什么（数据类型）
func main() {

	a := 10
	fmt.Printf("%T\n", a)
	b := [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", b)
	c := true
	fmt.Printf("%T\n", c)

	//	函数的类型
	wa1()
	fmt.Printf("%T\n", wa1)

	//wa2()
	fmt.Printf("%T\n", wa2)
	//	函数在Go语言中本身也是一个数据类型，加了（） 是调用函数 ， 不加（） 函数也是一个变量 可以赋值给别人

	//函数的类型就等于该函数创建的类型，也可以赋值给其他相同的函数
	var wa3 func(int, int, ...string) (int, int)
	wa3 = wa2
	q, w := wa3(1, 3, "S")
	fmt.Println(q, w)

}

// 无参无返回值
func wa1() {

}

// 有参有返回值
func wa2(a, b int, c ...string) (int, int) {
	return 0, 0
}
