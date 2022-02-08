package main

import "fmt"

func main() {
	//var n int = 20
	////演示goto的使用
	//fmt.Println("ok1")
	//fmt.Println("ok2")
	//if n > 10 {
	//	goto label1
	//}
	//fmt.Println("ok3")
	//label1:
	//fmt.Println("ok4")
	//fmt.Println("ok5")
	//fmt.Println("ok6")

	var n int = 20
	//演示return的使用
	fmt.Println("ok1")
	fmt.Println("ok2")
	if n > 10 {
		return
	}
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
	fmt.Println("ok6")


}