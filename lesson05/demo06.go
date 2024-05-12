package main

import "fmt"

// map的使用
func main() {
	var map1 map[int]string
	map1 = make(map[int]string)
	// 在map中如果你的key重复了，它就会覆盖这个key原来的值，一个key只能对应一个value
	map1[100] = "Go"
	map1[100] = "Go2"
	map1[500] = "java"

	fmt.Println(map1)
	fmt.Println(map1[100])
	fmt.Println(map1[1]) // 不存在，默认值 string ""

	// map中，没有index下标，有时候我们取值就需要判断这个key是否存在了
	// map中的判断，ok-idiom 是否存在的
	// value = map[key] , 隐藏的返回值 ok-idiom , 可选参数
	value, ok := map1[1]
	if ok {
		fmt.Println("map key 存在， value:", value)
	} else {
		fmt.Println("map key 不存在")
	}
	// 如果数据存在，修改它，如果不存在，就会创建对应的 key-value
	map1[100] = "aaa"
	fmt.Println(map1)
	map1[1] = "xxxx"
	fmt.Println(map1)

	//	map中删除数据
	delete(map1, 100)
	fmt.Println(map1)
	//	map的大小
	fmt.Println(len(map1))

}
