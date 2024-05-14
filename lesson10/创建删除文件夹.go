package main

import (
	"fmt"
	"os"
)

// 创建目录
// 项目开源框架，一运行，就会自动生成脚手架目录
func main() {
	// 打开一个文件夹（1、存在我就打开  2、不存在，创建这个文件夹）
	// func Mkdir(name string, perm FileMode) error
	// ModePerm : 0777
	err := os.Mkdir("D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\file1", os.ModePerm)
	if err != nil {
		//  存在就无法创建了 Cannot create a file when that file already exists.
		fmt.Println(err)
		//return
	}
	fmt.Println("创建文件夹完毕")
	// 创建层级文件夹
	err2 := os.MkdirAll("D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\file2\\aa\\bb\\cc\\dd", os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("层级文件夹创建完毕")

	// 删除 remove
	// func Remove(name string) error
	// 通过remove方法只能删除单个空文件夹：
	// remove D:\Environment\GoWorks\src\xuego\lesson11\file2: The directory is not empty.
	err3 := os.Remove("D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\file2\\aa\\bb\\cc\\dd")
	if err3 != nil {
		fmt.Println(err3)
		//return
	}
	fmt.Println("file delete success！！")
	// 如果存在多层文件，removeAll，相对来说比较危险，删除这个目录下的所有东西, 强制删除
	err4 := os.RemoveAll("D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\file2\\aa\\bb")
	if err4 != nil {
		fmt.Println(err4)
		return
	}
	fmt.Println("err4 delete success！！")
}
