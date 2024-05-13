package main

import "fmt"

// 方法重写，建立在父类和子类结构上的
type Animal struct {
	name string
	age  int
}

func (animal Animal) eat() {
	fmt.Println(animal.name, "正在吃...")
}
func (animal Animal) sleep() {
	fmt.Println(animal.name, "正在sleep...")
}

// 子类
type Dog2 struct {
	Animal
}

func (dog2 Dog2) w() {
	fmt.Println("waw~~~~~~~~~~")
}

type Cat2 struct {
	Animal
	color string //子类可以定义自己的属性
}

// 子类重写父类的方法 子类的方法名和父类同名，即可重写父类
func (cat2 Cat2) eat() {
	fmt.Println(cat2.name, "cat正在吃")
}

func main() {
	//定义一个子类使用父类方法
	dog2 := Dog2{Animal{"旺财", 3}}
	dog2.eat()
	//调用自己扩展的方法
	dog2.w()
	cat2 := Cat2{Animal{"咪咪咪", 2}, "黑色"}
	cat2.eat()
	fmt.Println(cat2.color)

	// 子类可以操作父类的方法，父类可以操作子类的吗？不可以
	// 父类不能调用子类自己的扩展的方法
	a := Animal{name: "大黄", age: 3}
	a.eat()
	a.sleep()
}
