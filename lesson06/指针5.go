package main

import "fmt"

// 指针函数
func main() {
	//调用函数后 得到指针类型的变量
	ptr := f1()
	fmt.Println("ptr:", ptr)
	fmt.Printf("ptr的类型:%T\n", ptr)
	fmt.Printf("ptr的地址:%p\n", &ptr)
	fmt.Printf("ptr的地址中的值:%d\n", *ptr)
	ptr[0] = 159
	fmt.Println(ptr[0])

	ptr2 := f3()
	ptr2[0] = 111
	fmt.Println("ptr2:", ptr2)
	fmt.Printf("ptr2的类型:%T\n", ptr2)
	fmt.Printf("ptr2的地址:%p\n", &ptr2)
	fmt.Printf("ptr2的地址中的值:%d\n", *ptr2)
}

// 返回一个指针数组
func f1() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	return &arr
}
func f2() *int {
	arr2 := 1
	return &arr2
}
func f3() *[4]int {
	arr3 := [4]int{999, 58, 6, 422}
	var p1 *[4]int = &arr3
	return p1
}
