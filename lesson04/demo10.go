package main

import "fmt"

// 切片扩容的内存分析
// 结论
// 1、每个切片引用了一个底层的数组
// 2、切片本身不存储任何数据，都是底层的数组来存储的，所以修改了切片也就是修改了这个数组中的数据
// 3、向切片中添加数据的时候，如果没有超过容量，直接添加，如果超过了这个容量，就会自动扩容，成倍的增加， copy
//   - 分析程序的原理
//   - 看源码
//
// 4、切片一旦扩容，就是重新指向一个新的底层数组。
func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 4, 5)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 0, 6, 1)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 10, 61, 21)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 4, 5)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 1, 5, 55, 81, 26, 17, 0)
	fmt.Println(s1)
	fmt.Printf("len:%d,cap%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

}
