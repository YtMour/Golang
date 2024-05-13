package main

// panic
func main() {

	// 什么时候会发生恐慌panic，我们不知道什么时候会报错
	// 程序运行的时候会发生panic

	// 手动抛出panic。我们在一些能够预知到危险的情况下，可以主动抛出
	// 如果有panic发生，我们尽可能接收它，并处理。
	// func panic(v any)  使用 panic() 程序就会终止,停在这里强制结束
	panic("程序 panic了")

}
