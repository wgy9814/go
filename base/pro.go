package main

import "fmt"

func main() {

	//基本数据转换
	var str = "hello " + "world"
	str += "wuguangyuan"
	fmt.Println(str)
	var a int
	var b float32
	var c float64
	fmt.Printf("a=%d,b=%f,c=%f",a,b,c)
	var n1 int32 = 12
	//var n2 int8
	var n3 int8
	n3 = int8(n1) + 128
	fmt.Printf("n3=%d",n3)
}