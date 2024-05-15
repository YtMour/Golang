package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	fmt.Println(Add[int](a, b))
	var a1 float64 = 1.55
	var a2 float64 = 2.22

	fmt.Println(Add[float64](a1, a2))

	// 每次都去写T的类型是很麻烦的，支持自动推导！
	// Go的泛型语法糖：自动推导 （本质，就是编译器帮我们加上去了，在实际运行，这里T还是加上去的）
	fmt.Println(Add(a, b))
}

// 真正的Add实现，传递不同的参数都是可以适配的！ Add[T]  T在调用的时候需要实例化
// 这种带了类型形参的函数就叫做泛型函数，极大的提高代码的灵活心，降低阅读性！
func Add[T int | float32 | float64](a T, b T) T {
	return a + b

}
