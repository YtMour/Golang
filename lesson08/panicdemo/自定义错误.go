package main

import "fmt"

type R struct {
	msg  string
	code int
	//	错误的信息可以定义多个 主要是描述这个错误出现的具体问题
	data []string
}

// (调用者) Error(参数)
func (r *R) Error() string {
	return fmt.Sprintf("%s,%d", r.msg, r.code)
}

// 错误自己的方法，提高一些操作效率
//func (e *R) ok() string {
//
//	return "success"
//}

// error的使用 如果函数或方法中有预期的错误，都需要抛出
func testErr(i int) (int, error) {
	if i != 0 {
		return 0, &R{msg: "错误数据", code: 500}
	}
	return i, nil
}

func main() {
	//假设这个方法返回了多个错误
	i, err := testErr(1)
	if err != nil {
		fmt.Println(err)
		e, ok := err.(*R)
		if ok {
			fmt.Println(e.msg)
			//	可以返回给前端的接口
		}
	}
	fmt.Println(i)
}
