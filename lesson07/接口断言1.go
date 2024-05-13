package main

import "fmt"

// 断言  t := i.(T)  t: t就是i接口是T类型的  i:接口   T:类型
// 语法  t,ok := i.(T)  ok 隐藏返回值 如果断言成功 ok就是true 否则就是 fales
func main() {
	//assertsString("1111a")
	//assertsString(true)
	assertInt(1)

}

// 判断一个变量是不是string类型的
func assertsString(i interface{}) {
	//如果断言失败 则会抛出 panic 恐慌，程序会停止执行
	s := i.(string)
	fmt.Println(s)
}

// 断言失败的情况
func assertInt(i any) {
	r, ok := i.(int)
	if ok {
		fmt.Println("是我们期望的结果 int")
		fmt.Println(r)
	} else {
		fmt.Println("不是我们期望的结果")
	}
}
