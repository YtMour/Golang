package main

import "fmt"

// 数组指针
func main() {

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("arr1", arr1)
	fmt.Printf("arr1的地址:%p\n", &arr1)
	//	创建指针 指向这个数组来进行操作
	var p1 *[4]int = &arr1
	fmt.Printf("p1指向的地址:%p\n", p1)
	fmt.Printf("p1自己的地址:%p\n", &p1)
	fmt.Printf("p1指向地址的值:%d\n", *p1)
	//	操作数组指针
	(*p1)[0] = 25
	fmt.Println("arr1", arr1)
	fmt.Printf("p1指向地址的值:%d\n", *p1)
	// 语法糖：由于p1指向了arr1这个数组，所以可以直接用p1来操控数组
	// 指向了谁，这个指针就可以代表谁。
	// p1 = arr1
	p1[1] = 195
	fmt.Println("arr1", arr1)
	fmt.Printf("p1指向地址的值:%d\n", *p1)
}
