package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	//	传输源文件地址
	srcFile := "C:\\Users\\root\\Desktop\\client\\头像.jpg"
	//	传输的目标地址

	destFile := "D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\server\\new.jpg"
	//	临时记录文件
	tempFile := "D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\temp.txt"

	//	创建对应的file对象
	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	//关闭
	defer file1.Close()
	defer file2.Close()
	fmt.Println("file1/2/3 文件连接建立完毕")
	//	读取temp
	file3.Seek(0, io.SeekStart)
	buf := make([]byte, 1024, 1024)
	n, _ := file3.Read(buf)
	//	转换成String - 数字
	countStr := string(buf[:n])
	//fmt.Println(countStr)
	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println("temp中记录的值为:", count)

	//	设置读写的偏移量
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	fmt.Println("file1/2 光标已经移动到了目标位置")

	//	开始读写(复制,上传)
	bufData := make([]byte, 1024, 1024)

	//	记录读取了多少个字节
	total := int(count)

	for {
		// 读取数据
		readNum, err := file1.Read(bufData)
		if err == io.EOF { // file1 读取完毕了
			fmt.Println("文件传输完毕了")
			file3.Close()
			os.Remove(tempFile)
			break
		}

		// 向目标文件中写入数据
		writeNum, err := file2.Write(bufData[:readNum])

		// 将写入数据放到 total中, 在这里total 就是传输的进度
		total = total + writeNum
		// temp.txt 存放临时记录数据
		file3.Seek(0, io.SeekStart) // 将光标重置到开头
		file3.WriteString(strconv.Itoa(total))
		fmt.Println(total)

		//	模拟断电
		if total > 50000 {
			panic("断电了")
		}

	}

}
