package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  string
}

func (user User) Say(msg string) {
	fmt.Println("User 说：", msg)
}

// toString : 打印结构体信息
func (user User) PrintInfo() {
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", user.Name, user.Age, user.Sex)
}

func main() {
	user := User{"kuangshen", 18, "男"}

	reflectGetInfo(user)
}

// 通过反射，获取变量的信息
func reflectGetInfo(v interface{}) {
	// 1、获取参数的类型Type , 可能是用户自己定义的，但是Kind一定是内部类型struct
	getType := reflect.TypeOf(v)
	fmt.Println(getType.Name()) // 类型信息 User
	fmt.Println(getType.Kind()) // 找到上级的种类Kind  struct

	// 2、获取值
	getValue := reflect.ValueOf(v)
	fmt.Println("获取到value", getValue)

	// 获取字段，通过Type扒出字段
	// Type.NumField() 获取这个类型中有几个字段  3
	// field(index) 得到字段的值
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)              // 类型
		value := getValue.Field(i).Interface() // value
		// 打印
		fmt.Printf("字段名：%s,字段类型：%s，字段值：%v\n", field.Name, field.Type, value)
	}

	// 获取这个结构体的方法 , 获取方法的数量
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法的名字：%s,方法类型：%s", method.Name, method.Type)
	}

}

/*
type User struct{
   Name string
   Age int
   Sex string
}
func (main.User) PrintInfo(){}
func (main.User) Say(string)
*/
