package main

import "fmt"

// 多态
// 定义接口
type Animal2 interface {
	eat()
	sleep()
}
type Dog3 struct {
	name string
}

func (dog3 Dog3) eat() {
	fmt.Println(dog3.name, "--eat")
}
func (dog3 Dog3) sleep() {
	fmt.Println(dog3.name, "--sleep")

}

func main() {
	//	Dog3 两重身份 ：1.Dog3  2.Animal2   多态
	dog3 := Dog3{"小黑"}
	dog3.eat()
	dog3.sleep()

	//	Dog 也可以是Animal
	test2(dog3)

	// 定义一个类型可以为接口类型的变量
	// 实际上所有实现类都可以赋值给这个对象
	var animal Animal2 // 模糊的 -- 具体化，将具体的实现类赋值给他，才有意义
	animal = dog3
	test2(animal)

	test2(animal)
}

// Animal接口
func test2(a Animal2) {
	a.eat()
	a.sleep()
}
