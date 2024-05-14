package main

import (
	"fmt"
	"strconv"
)

// string convert = strconv
func main() {
	//bool
	s1 := "true"
	//	转换 - 字符串转bool(解析 parse)     bool转字符串(格式化format)
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1)
	if b1 == true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	//	bool转字符串
	s2 := strconv.FormatBool(b1)
	fmt.Printf("%T,%s\n", s2, s2)
	//	整数-字符串
	s3 := "66"
	//整数： 数字，进制，大小
	i1, err := strconv.ParseInt(s3, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i1, i1)
	s4 := strconv.FormatInt(i1, 10)
	fmt.Printf("%T,%s\n", s4, s4)

	// 10进制转换字符串，简便方法  atoi  itoa
	atoi, _ := strconv.Atoi("-20")
	fmt.Printf("%T,%d\n", atoi, atoi)
	itoa := strconv.Itoa(30)
	fmt.Printf("%T,%s\n", itoa, itoa)

}
