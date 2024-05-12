package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		fmt.Println(fb(i))
	}

}

func fb(n int) int {
	//结束递归的条件
	if n <= 2 {
		return 1
	} else {
		return fb(n-1) + fb(n-2)
	}
}
