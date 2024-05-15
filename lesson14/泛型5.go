package main

import "fmt"

type MySlice[T int | float32] []T

func main() {
	var s MySlice[int] = []int{1, 2, 3, 4}
	fmt.Println(s.sum())

	var s1 MySlice[float32] = []float32{1.7, 2, 3.5, 4}
	fmt.Println(s1.sum())
}

// 调用者，类型是不确定的，用户传什么，她就实例化什么。  类型参数化了 ， 泛型
// 没有泛型之前， 反射: 	reflect.ValueOf().Kind() , 也需要很多if，本质是逻辑相同的，只是类型不同！
func (s MySlice[T]) sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
