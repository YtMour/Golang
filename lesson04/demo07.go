package main

import "fmt"

// 切片
func main() {

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	var s1 []int //变长，长度是可变的
	fmt.Println(s1)

	if s1 == nil {
		fmt.Println("切片为空")
	}

	s2 := []int{1, 2, 3, 4} //切片 边长
	fmt.Println(s2)
	fmt.Printf("%T,%T\n", arr, s2)
	fmt.Println(s2[1])

}
