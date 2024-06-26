package main

import "fmt"

/*
	遍历map
	 - key、value 无序的，遍历map，可能每次的结果排序都不一致。
	  - "aaa" "aaaaa"
	  - "bbb" "bbbbb"
	  - "ccc" "ccccc"

1、map是无序的，每次打印出来的map可能都不一样，它不能通过index获取，只能通过key来获取
2、map的长度是不固定的，是引用类型的
3、len可以用于map查看map中数据的数量，但是cap无法使用
4、map的key 可以是 布尔类型，整数，浮点数，字符串
*/
func main() {
	map1 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	// 循环遍历，只能通过for range
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}
