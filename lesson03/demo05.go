package main

import "fmt"

func main() {
	fmt.Println(getSum(1, 2, 3, 4, 5, 6, 7, 8))
}

// 可变参数： 一个函数的参数类型确定，参数的个数不确定，可以使用可变参数
// 可变参数如果有多个参数必须放在最后一个参数
func getSum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

	}
	return sum
	// 接收多个参数 nums 可变参数
	// 使用下标来接收，下标是从0开始的
	// nums : 100,200,300,400,500
	// nums[0] = 100
	// nums[1] = 200
	// nums[2] = 300

	// 了解传递了多少个数字  len()函数，获取可变参数的长度
}
