package main

import (
	"fmt"
	"reflect"
)

// 反射
/*
Type : reflect.TypeOf(a) , 获取变量的类型
Value ：reflect.ValueOf(a) ， 获取变量的值
*/

func main() {
	var a float64 = 3.14
	fmt.Println("type:", reflect.TypeOf(a))
	fmt.Println("value:", reflect.ValueOf(a))

	// 根据反射的值，来获取对象对应的类型和数值
	// 如果我们不知道这个对象的信息，我们可以通过这个对象拿到代码中的一切。
	// Value
	v := reflect.ValueOf(a) // string int User
	// Kind : 获取这个值的种类， 在反射中，所有数据类型判断都是使用种类。
	if v.Kind() == reflect.Float64 {
		fmt.Println(v.Float())
	}
	if v.Kind() == reflect.Int {
		fmt.Println(v.Int())
	}
	fmt.Println(v.Kind() == reflect.Float64)
	//fmt.Println(v.Type())
}
