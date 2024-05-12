package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		for i := 0; i < 5; i++ {
			fmt.Print("*")

		}
		fmt.Println()
	}

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d   ", j, i, i*j)
		}
		fmt.Println()

	}
	//==========================================================
	//1、打印上面那个三角形
	for i := 1; i <= 6; i++ {
		// 如何递减实现空格 5  4  3  2  1
		for j := 1; j <= 6-i; j++ {
			fmt.Print(" ")
		}
		// i         1 2 3 4 5 6
		// 实现三角形 1 3 5 7 9 11   2i-1
		// 第一行  *    需要循环1次
		// 第二行 ***   需要循环3次
		// 第三行 ***   需要循环5次
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		// 换行
		fmt.Println()
	}
	// 第二个三角形，倒三角

	// 5 4 3 2 1 0
	for i := 5; i >= 1; i-- {
		// 如何递增实现空格 1 2 3 4 5
		for j := 1; j <= 6-i; j++ {
			fmt.Print(" ")
		}
		// 实现三角形 1 3 5 7 9 11   2i-1
		// 第一行  9    需要循环9次
		// 第二行  7   需要循环7次
		// 第三行  5   需要循环5次
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
