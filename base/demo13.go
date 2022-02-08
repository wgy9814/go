package main

import (
	"fmt"
	"math"
)

func main() {
	//if条件表达式不能是赋值语句，编译不通过
	//单分支
	//声明两个float64型变量并赋值，判断第一个数大于10.0 并且 第二个数小于20.0，打印两数之和
	var n3 float64 = 11.0
	var n4 float64 = 18.0
	if n3 > 10.0 && n4 < 20.0 {
		fmt.Println(n3 + n4)
	}

	//声明两个int32型变量并赋值，判断两者的和，是否能被3又能被5整除，打印提示信息

	var num1 int32 = 5
	var num2 int32 = 10
	if (num1 + num2) % 3 == 0 && (num1 + num2) % 5 == 0 {
		fmt.Println("能被3又能被5整除")
	}

	//判断一个年份是否闰年，闰年的符合条件是符合一下两者之一即可
	//1年份能被4整除，但不能被100整除；2能被400整除
	var year int = 2020
	if (year % 4 == 0 && year % 100 != 0) ||  year % 400 == 0 {
		fmt.Println("是闰年")
	}

	//多分支
	//求ax2+bx+c=0方程的根，a,b,c分别为函数的参数，如果：b2-4ac>0,则有两个解
	//b2-4ac=0,则有一个解：b2-4ac<0,则无解
	//提示1：x1 = (-b+sqrt(b2-4ac))/2a
	//		x2 = (-b-sqrt(b2-4ac))/2a
	//提示2：math.Sqrt(num);求平方根 需要引入math包

	//分析思路
	//1. a b c 是3个float64
	//2. 使用到给出数学公式
	//3. 使用到多分支
	//4. 使用math.Sqrt方法 =》手册

	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 6.0

	m := b * b - 4 * a * c

	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v,x2=%v",x1,x2)
	}else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v",x1)
	}else {
		fmt.Println("无解")
	}




}