package main

import (
	"fmt"
	// 匿名导入包，会执行报下所有go文件的 init 函数, 单个init被多个地方导入，只会执行一次
	// 1、先执行导入包的init函数，单个go文件中是顺序执行的，所有go中的init函数执行完毕后，才会到main包
	// 2、如果导入了多个匿名包，按照main中导入包的顺序来进行执行。
	// 3、在同一个包下的go文件如果有多个，都有init的情况下，按照文件排放顺序来执行对应的init函数（）
	_ "MyGodemo/lesson08/test"
	_ "MyGodemo/lesson08/zzz"
)

func init() {
	fmt.Println("main ---init 1")
}

func main() {
	// init 函数不需要传入参数，也没有返回值，任何地方不能调用 init()
	//init()
}
