package main

import "fmt"

// 方法：可以理解为函数多了一个调用者
// 方法可以重名，不同的对象，调用的结果是不一样的
type Dog struct {
	name string
	age  int
}

// 方法定义,   func 方法调用者  方法名()
// 1、方法可以重名，只需要调用者不同
// 2、如果调用者相同，则不能重名
func (dog Dog) eat() {
	fmt.Println("dog eat")
}
func (dog Dog) sleep() {
	fmt.Println("dog sleep")
}

// ===================
type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("cat eat")
}
func (cat Cat) sleep() {
	fmt.Println("cat sleep")

}

func main() {
	//	创建一个对象
	dog := Dog{"旺财", 2}
	fmt.Println(dog)
	//	方法的调用
	dog.eat()

	cat := Cat{"咪咪", 1}
	fmt.Println(cat)
	cat.sleep()
}
