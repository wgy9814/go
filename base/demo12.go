package main

import "fmt"

func main() {
	//单分支
	//编写一个程序，可以输入年龄，如果大于18岁，则输出“你年龄大于18岁了，该对自己负责了”

	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你年龄大于18岁了，该对自己负责了")
	}

	//golang支持在if中，直接定义一个变量，比如下面
	//if age := 20; age > 18 {
	//	fmt.Println("你年龄大于18岁了，该对自己负责了")
	//}

	//双分支
	//编写一个程序，可以输入年龄，如果大于18岁，则输出“你年龄大于18岁了，该对自己负责了”，否则输出“你的年龄不大这次就放过你了”
	var ages int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&ages)
	if ages > 18 {
		fmt.Println("你年龄大于18岁了，该对自己负责了")
	} else {
		fmt.Println("你的年龄不大这次就放过你了")

	}

}