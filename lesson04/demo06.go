package main

import (
	"fmt"
)

// 多维
func main() {
	arr := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println(arr[0][1])

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Println(arr[i][j])
		}
		//fmt.Println(arr[i])
	}

	for Index, value := range arr {
		fmt.Println(Index, value)

	}
}
