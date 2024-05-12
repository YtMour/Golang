package main

import "fmt"

func main() {
	//返回多个返回值
	//交换两个string
	a, b := swap("yt", "aaa")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}
