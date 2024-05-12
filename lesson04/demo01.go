package main

import "fmt"

// 数组
func main() {
	// array数组定义，变量
	// 数组也是一个数据类型
	// 数组的定义：  [数组的大小size]变量的类型 ，
	// 我们定义了一组这个类型的数组集合，大小为size，最多可以保存size个数
	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 20
	arr1[2] = 35
	arr1[3] = 77
	arr1[4] = 100
	fmt.Println(arr1)
	fmt.Println(arr1[1])

	//	数组中的常用方法  len()取长度 cap() 容量
	fmt.Println("数组的长度", len(arr1))
	fmt.Println("数组的容留", cap(arr1))

	arr1[1] = 159
	fmt.Println(arr1)
	fmt.Println(arr1[1])
}
