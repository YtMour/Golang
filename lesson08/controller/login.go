package controller

// 包
//
//import "math/rand"  // 随机数生成

import (
	//"crypto/rand"
	//R "math/rand" // 可以给包起别名
	//. "math/rand" // 简便模式：可以直接调用该包下的函数，不需要通过包名。
	_ "math/rand" // 匿名导入，只会执行这个包下的init方法
)

func test() {

}
