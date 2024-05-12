package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Println("i=", i, "sum=", sum)
		sum += i
	}
	fmt.Println(sum)

	for i := 1; i <= 10; i++ {
		if i == 5 {
			//break
			continue // 到这里就结束了当次循环，不会向下了，继续从头开始
		}
		fmt.Println(i)
	}
}
