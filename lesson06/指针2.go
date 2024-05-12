package main

import "fmt"

// 指针的使用
/*
1、定义指针变量
2、为指针变量赋值  &
3、访问指针变量中地址所指向的值   *
* ：在指针类型前面加上 * 号，就是来获取指针所指向的地址的值

*/
func main() {
	var a int = 10
	fmt.Printf("a 变量的值：%d\n", a)
	fmt.Printf("a 变量的地址：%d\n", &a)
	// 声明 指针变量，指向a， 指针其实就是一个特殊的变量而已。，ptr命名  p
	// 定义变量格式  var ptr *类型
	var p *int
	p = &a
	//var p2 = &a

	fmt.Printf("p 变量存储的指针地址：%p\n", p)
	fmt.Printf("p 变量自己的地址：%p\n", &p)
	fmt.Printf("p 变量存储的指针地址指向的值：%d\n", *p)

	//	改变a的值
	a = 20
	fmt.Printf("a:%d\n", a)
	fmt.Printf("*p:%d\n", *p)
	//通过指针改变a
	*p = 30
	fmt.Printf("a:%d\n", a)
	fmt.Printf("*p:%d\n", *p)
	//	指针的套娃
	var ptr **int
	ptr = &p
	fmt.Printf("ptr变量存储的指针的地址：%p\n", ptr)
	fmt.Printf("ptr变量自己的地址：%p\n", &ptr)
	fmt.Printf("*ptr变量存储的地址：%p\n", *ptr)
	fmt.Printf("*ptr变量存储的地址中的值：%d\n", **ptr)
	**ptr = 159
	fmt.Println(a)

	var ptr2 ***int
	ptr2 = &ptr
	***ptr2 = 963
	fmt.Println(a)
	var ptr3 ****int
	ptr3 = &ptr2
	****ptr3 = 984516854
	fmt.Println(a)
	//	逆向 加解密  壳：
}
