package main

import "fmt"

// 自定义错误
type YtError struct {
	code int
	msg  string
}

// 处理error的逻辑
func (e YtError) print() bool {
	fmt.Println("hhee")
	return true
}

func (Y *YtError) Error() string {
	return fmt.Sprintf("错误信息: %s,错误代码：%d\n", Y.msg, Y.code)
}

// 使用错误测试
func tt(i int) (int, error) {
	//需要编写的错误
	if i != 0 {
		// 更多的时候我们会使用自定义类型的错误
		return i, &YtError{msg: "非预期数据", code: 500}
	}
	//正常结果
	return i, nil
}
func main() {
	i, err := tt(3)
	if err != nil {
		fmt.Println(err)
		Yt_err, ok := err.(*YtError)
		if ok {
			if Yt_err.print() {
				//处理这个错误中的子错误逻辑
			}
			fmt.Println(Yt_err.msg)
			fmt.Println(Yt_err.code)
		}
	}
	fmt.Println(i)
}
