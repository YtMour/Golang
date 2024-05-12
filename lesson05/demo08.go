package main

import "fmt"

// map结合slice使用
/*
需求：
1、创建map来存储人的信息，name，age,sex，addr
2、每个map保存一个的信息
3、将这些map存入到切片中
4、打印这些数据
*/
func main() {
	u1 := make(map[string]string)
	u1["name"] = "Yt"
	u1["age"] = "27"
	u1["sex"] = "男"
	u1["addr"] = "hh"

	u2 := map[string]string{"name": "Yt2", "age": "21", "sex": "女", "addr": "北京"}
	fmt.Println(u2)
	//	存放到切片中
	uDatas := make([]map[string]string, 0, 2)
	uDatas = append(uDatas, u1)
	uDatas = append(uDatas, u2)
	fmt.Println(uDatas)
	for _, v := range uDatas {
		fmt.Println(v["name"], v["age"])
	}
}
