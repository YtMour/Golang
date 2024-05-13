package main

import "fmt"

/*
Go语言中结构体嵌套
1、模拟继承   is - a

	type A struct{
	   field
	}
	type B struct{
	   A  // 匿名字段  B.A
	   field
	}
	// B是可以直接访问A的属性的

2、聚合关系   has - a

	type C struct{
	   field
	}
	type D struct{
	   c C  //聚合关系
	   field
	}
	// D是无法直接访问C中的属性的，必须使用D.c.c_field
*/
type Person struct {
	name string
	age  int
}

// 定义一个子类  拥有父类所有的属性
type Student struct {
	school string
	Person //匿名字段 实现继承
}

func main() {
	//创建父类对象
	p1 := Person{"James", 20}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)
	//	创建子类
	s1 := Student{"sanxia", Person{"ppp", 18}}
	fmt.Println(s1)
	fmt.Println(s1.Person.name, s1.Person.age)

	var s2 = Student{"北大", Person{"zhangsan", 18}}
	fmt.Println(s2)
	// 概念：提升字段, 只有匿名字段才可以做到
	// Person 在Student中是一个匿名字段， Person中的属性 name age 就是提升字段
	// 所有的提升字段就可以直接使用了，不同在通过匿名字段来点了
	var s3 Student
	s3.name = "lishi"
	s3.age = 12
	s3.school = "清华"
	fmt.Println(s3)
	fmt.Println(s3.name, s3.age, s3.school)
}
