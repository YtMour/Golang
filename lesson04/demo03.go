package main

import "fmt"

func main() {
	arr1 := [5]int{10, 21, 32, 40, 59}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	fmt.Println("============")
	// goland 快捷方式 数组.for，未来循环数组、切片很多时候都使用for    range
	// for 下标,下标对应的值  range 目标数组切片
	// 就是将数组进行自动迭代。返回两个值 index、value
	// 注意点，如果只接收一个值，这个时候返回的是数组的下标
	// 注意点，如果只接收两个值，这个时候返回的是数组的下标和下标对应的值
	for index, value := range arr1 {
		fmt.Println("index:", index, "value:", value)
	}
	for _, value := range arr1 {
		fmt.Println(value)
	}

}
