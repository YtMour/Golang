package main

import "fmt"

func main() {
	r1, r2 := fun1(2, 4)
	fmt.Println(r1, r2)
}
func fun1(len, wid float64) (zc, area float64) {
	area = len * wid
	zc = (len + wid) * 2
	fmt.Println("zc:", zc)
	fmt.Println("area:", area)
	return
}
