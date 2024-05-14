package main

import (
	"fmt"
	"os"
)

func main() {
	//	创建文件 Create
	file1, err := os.Create("lesson10/file3/a.go") // 相对路径
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file1)
	// 删除
	os.Remove("D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\file3\\a.go")
}
