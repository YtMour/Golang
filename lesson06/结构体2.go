package main

import "fmt"

type User2 struct {
	name string
	age  int
	sex  string
}

func main() {
	user1 := User2{"YT", 18, "男"}
	fmt.Println(user1)
	//结构体类型  包 . struc
	fmt.Printf("%T,%p\n", user1, &user1)
	//	结构体是值类型还是引用类型
	user2 := user1
	fmt.Println(user2)
	//结构体类型  包 . struc
	fmt.Printf("%T,%p\n", user2, &user2)
	user2.name = "qwrqwrqwrq"
	fmt.Println(user2)
	//	指针解决值传递的问题
	var user_ptr *User2
	user_ptr = &user1
	fmt.Println(*user_ptr)
	fmt.Printf("%T,%p\n", user_ptr, &user_ptr)
	user_ptr.name = "scv255555"
	fmt.Println(*user_ptr)

	//	内置函数 new 创建对象  new关键字创建的对象，都返回指针 而不是对象
	// func new(Type) *Type
	// 通过这种方式创建的结构体对象更加灵活，突破了结构体是值类型的限制。
	user3 := new(User2)
	fmt.Println(*user3)
	user3.name = "oott"
	fmt.Println(*user3)
	updateUser(user3)
	fmt.Println(user3)

}
func updateUser(user *User2) {
	user.age = 18
}
