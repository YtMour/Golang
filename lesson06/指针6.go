package main

import "fmt"

func main() {

	a := 10
	fmt.Println("改变前a：", a)
	fmt.Println("改变前a的地址：", &a)

	a1(&a)
	fmt.Println("改变后a：", a)
	fmt.Println("改变后a的地址：", &a)

}

// 指针当作函数的参数
func a1(ptr *int) {
	fmt.Println("ptr：", ptr)
	fmt.Println("ptr的地址：", &ptr)
	fmt.Println("ptr指向的地址的值：", *ptr)
	*ptr = 20
}

func a2(a int) {
	fmt.Println("改变前a：", a)
	a = 20
	fmt.Println("改变后a：", a)

}
