package main

import (
	"fmt"
	"strconv"
)
//string转成基本数据使用
func main() {
	var str string = "true"
	var b bool
	//b,_ = strconv.ParseBool(str)
	//说明
	// 1. strconv.ParseBool(str) 函数会返回两个值（value bool,err error）
	// 2. 只想要value bool 不想要err 所以用_忽略

	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b =%v\n",b,b)

	var str2 string = "1234590"
	var n1 int64
	var n2 int
	n1 , _ = strconv.ParseInt(str2,10,64)
	n2 = int(n1)
	fmt.Printf("n1 type %T b =%v\n",n1,n1)
	fmt.Printf("n2 type %T b =%v\n",n2,n2)

	var str3 string = "123.456"
	var f1 float64
	f1 , _ = strconv.ParseFloat(str3,64)
	fmt.Printf("f1 type %T b =%v\n",f1,f1)

	//注意没转成功就默认值
	var str4 string = "hello"
	var n3 int64 = 11
	n3 , _ = strconv.ParseInt(str4,10,64)
	fmt.Printf("n3 type %T n3 =%v\n",n3,n3)


}