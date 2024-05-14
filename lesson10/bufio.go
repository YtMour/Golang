package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// bufio 的应用
func main() {
	file, err := os.Open("D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\测试文本.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// 读取文件
	//创建一个bufio包下的 reader对象。
	bufioReader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	n, err := bufioReader.Read(buf)
	fmt.Println("读取到了多少个字节：", n)
	fmt.Println("读取的内容:", string(buf[:n]))
	//
	//// 读取键盘的输入
	//// 键盘的输入，实际上是流 os.Stdin
	inputReader := bufio.NewReader(os.Stdin)
	// delim 到哪里结束读取
	readString, _ := inputReader.ReadString('\n')
	fmt.Println("读取键盘输入的信息：", readString)

}
