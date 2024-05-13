package main

import (
	"errors"
	"fmt"
)

// 自己定义一个错误
// 1、errors.New("xxxxx")
// 2、fmt.Errorf()
// 都会返回  error 对象， 本身也是一个类型
func main() {
	age_err := setAge(-1)
	if age_err != nil {
		fmt.Println(age_err)
	}
	fmt.Printf("%T\n", age_err)

	//	方法二
	errInfo1 := fmt.Errorf("我是一个错误信息:%d\n", 500)
	errInfo2 := fmt.Errorf("我是一个错误信息:%d\n", 404)
	if errInfo1 != nil {
		fmt.Println(errInfo1)
		fmt.Println(errInfo2)
	}
}

// 设置年龄的函数
func setAge(age int) error {
	if age < 0 {
		//抛出一个错误
		//age = -1
		return errors.New("年龄不合法")
		//给出一个默认值
	}
	if age > 200 {
		return errors.New("年龄不合法")
	}
	fmt.Println("年龄设置成功：age=", age)
	return nil
}
