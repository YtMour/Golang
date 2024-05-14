package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "lesson10/测试文本.txt"
	// 权限：如果我们要向一个文件中追加内容， O_APPEND, 如果没有，就是从头开始写
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_RDONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()
	//	操作
	bs := []byte{65, 66, 67, 68, 69, 70}
	n, err := file.Write(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	// string类型的写入
	n, err = file.WriteString("hhahahahah哈哈哈哈哈哈哈")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

}
