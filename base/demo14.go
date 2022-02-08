package main

import "fmt"

func main() {
	//参加百米运动会，如果用时 8 秒以内进入决赛，否则提示淘汰。并且根据性别提示进入男子组或女子组。输入成绩和性别
	//输入成绩和性别
	//分析思路
	//1. 定义一个变量，跑步秒数，float64
	//2. 定义一个变量，性别，string
	//3. 嵌套分支
	var second float64

	fmt.Println("请输入秒数")
	fmt.Scanln(&second)
	if second <= 8 {
		//进入决赛
		var gender string
		fmt.Println("请输入性别")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入男子组")
		} else {
			fmt.Println("进入女子组")
		}

	}else{
		fmt.Println("out.....")
	}





}