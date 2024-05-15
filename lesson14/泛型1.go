package main

import (
	"fmt"
)

func main() {
	is := []int{1, 2}
	//strs := []string{"kuangshen", "xuexiangban"}
	printSlice2(is)
	//printSlice(strs)
}

// 如何实现一个方法可以打印上面不同类型的切片呢？  反射
func printSlice(s interface{}) {
	// 断言 x.(T), 如果x实现了T，那么就将 x转换为T类型
	for _, v := range s.([]string) {
		fmt.Println(v)
	}
}

// 参数的类型是不确定的，让用户自己指定。
// 泛型也是使用 []，和数组很像！
func printSlice2[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
