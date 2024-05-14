package main

import "MyGodemo/lesson10/copy"

func main() {
	source := "C:\\Users\\root\\Desktop\\头像.jpg"
	dest := "D:\\IDEAGoLang\\GoCod\\GoWorks\\src\\MyGodemo\\lesson10\\头像-copy.jpg"
	copy.Copy2(source, dest)

}
