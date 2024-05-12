package main

import "fmt"

// 数组是值类型  所有的赋值后的对象修改值后不影响原来的对象。
func main() {

	//	数组类型的样子 %T
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)

	arr3 := arr1
	fmt.Printf("%T\n", arr3)
	fmt.Println(arr1)
	fmt.Println(arr3)
	arr3[0] = 15
	fmt.Println(arr3)

}
