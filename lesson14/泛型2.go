package main

import "fmt"

// s1 是我们自己定义的类型 ，本质 []int
type s1 []int
type s2 []float64
type s3 []float32

// 我们定义的结构都是一样的，只是它的类型不同，就需要重新定义这么多的类型。
// 思考：是否有一种机制，只定义一个类型就可以代表上面的所有类型？
// 泛型：类型 参数化了！   参数：人为传递的

/*
1、T 说白了就是一个占位符，类型的形式参数，T是不确定的，需要在使用的时候进行传递。
2、由于T类型是不确定的，我们需要加一些约束 int|float64|float32 。告诉编译器我这个T，只接受
  int、float64、float32 类型
3、我们这里定义的类型是什么？Slice[T]
*/

// 这种类型的定义方式，带了类型形参，和普通定义类型就完全不同的。
// 普通的定义类型，这个类型只能代表本身一个，泛型类型，我们可以实现，参数类型传递。
// 我们可以在使用的时候来定义类型。
// 语法糖：简化开发
type Slice[T int | float64 | float32] []T

func main() {

	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("%T\n", a) // Slice[int]

	var b Slice[float64] = []float64{1.0, 2.0, 3.0}
	fmt.Printf("%T\n", b) // Slice[float64]

	// 不能够赋值（string 不在T的约束当中，不能实例化的）
	//var c Slice[string] = []string{"kuangshen","xxx"}

	// T是占位符，在使用的时候，必须要实例化为具体的类型。
	//var d Slice[T] = []int{1,2,3}
}

func test(name interface{}) {
	fmt.Println(name)
}
