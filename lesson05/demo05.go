package main

import "fmt"

// map 集合 保存数据的一种结构
func main() {
	// 创建一个map，也是一个变量，数据类型是 map
	// map[key]value
	var map1 map[int]string // 只是声明了但是没有初始化,不能使用 nil
	if map1 == nil {
		fmt.Println("map1==nil")
	}
	//	更多  使用make方法创建
	map2 := make(map[string]string)
	fmt.Println(map2)

	// 在创建的时候，添加一些基础数据
	// map[string]int nil
	// map[string]int {key:value,key:value,......}
	map3 := map[string]int{"GO": 100, "java": 50, "c": 10}
	fmt.Println(map3)
	// 关于map的类型，就如定义的一般 map[string]int
	// 类型主要是传参要确定
	fmt.Printf("%T\n", map3)

}
