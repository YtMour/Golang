package main

import (
	"fmt"
	"net/http"
)

// 手动实现一个请求响应 。 net/http
// http://127.0.0.1:8080/hello
func main() {
	// 访问这个url就会触发 hello 函数 （Request） ResponseWriter
	http.HandleFunc("/hello", hello)
	fmt.Println("项目已启动：127.0.0.1:8080")
	fmt.Println("访问hello请求测试：127.0.0.1:8080/hello")
	// 127.0.0.1 本机  localhost , 监听程序
	http.ListenAndServe("127.0.0.1:8080", nil)
}

// 接收的请求 Request
// ResponseWriter 响应
func hello(resp http.ResponseWriter, req *http.Request) {
	//
	fmt.Println("已收到来自浏览器的请求")
	fmt.Println("请求的地址:", req.URL)
	fmt.Println("请求的方法:", req.Method)

	// 给浏览器响应数据
	resp.Write([]byte("<h1>感谢大家选择学相伴Go语言全栈教学</h1>"))

}
