package main

import "fmt"

// 定一个结构体 type User struct
type User struct {
	name string
	age  int
	sex  string
}

func main() {
	// 通过结构体创建对象， 以前的类型都是用的基本类型，自己定义类型了type，Struct
	// 定义了结构体对象，不赋值，默认都是这个结构体的零值 {"",0,""}
	var user1 User
	user1.name = "Yt"
	user1.age = 20
	user1.sex = "男"
	fmt.Println(user1)
	fmt.Println(user1.name)

	// 创建对象的方式二
	user2 := User{}
	user2.name = "qwe"
	user2.age = 20
	user2.sex = "女"
	fmt.Println(user2)
	fmt.Println(user2.name)

	// 创建对象的方式三
	user3 := User{
		name: "feige",
		age:  35,
		sex:  "男",
	}
	fmt.Println("user:", user3)
	fmt.Println("user:", user3.name)
	// 创建对象的方式四，这种不声明属性的方式，需要参数一一匹配
	user4 := User{"guoguo", 18, "女"}
	fmt.Println("user:", user4)

}
