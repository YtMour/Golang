package main

import "fmt"

func main() {

	s1 := make([]int, 0, 5)
	fmt.Println(s1)
	//切片扩容   append()  切片是会自动扩容的。
	s1 = append(s1, 10, 10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println(s1)

	s2 := []int{100, 159, 551, 222}
	s1 = append(s1, s2...)
	fmt.Println(s1)
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}
}
