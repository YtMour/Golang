package main

import "fmt"

// 指针数组
func main() {
	a := 1
	b := 2
	c := 3
	d := 4
	//创建指针数组
	arr1 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr1)

	*arr1[1] = 100
	fmt.Println(arr1)
	fmt.Println(a, b, c, d)
	a = 159
	fmt.Println(*arr1[0])

}
