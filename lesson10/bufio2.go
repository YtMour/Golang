package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.OpenFile("D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\a.txt",
		os.O_APPEND|os.O_CREATE,
		os.ModePerm)
	defer file.Close()
	//	bufio
	fileWrite := bufio.NewWriter(file)
	writeNum, _ := fileWrite.WriteString("123595")
	fmt.Println("writeNum:", writeNum)
	// 发现并没有写出到文件，是留在了缓冲区，所以我们需要调用 flush 刷新缓冲区
	// 手动刷新进文件
	fileWrite.Flush()
}
