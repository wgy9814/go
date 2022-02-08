package main

import (
	"fmt"
	"go_code/chapter/demo04/model"
	"strconv"
)
//基本数据转成string使用
func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string

	//第一种方式fmt.Sprintf
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str =%v\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str =%v\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str =%v\n",str,str)

	str = fmt.Sprintf("%c",mychar)
	fmt.Printf("str type %T str =%q\n",str,str)

	//第二种方式 strconv
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = false

	str = strconv.FormatInt(int64(num3),10)
	fmt.Printf("str type %T str =%q\n",str,str)

	str = strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type %T str =%q\n",str,str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str =%q\n",str,str)

	//strconv包中有一个函数Itoa


	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str =%q\n",str,str)
	fmt.Println(model.HeroName)

	const (
		e = iota
		f
		g
	)
	fmt.Println(e, f, g)

}