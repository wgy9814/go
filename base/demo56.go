package main

import "fmt"

func main() {
	//第一种使用方式
	//map的声明和注意事项
	var a map[string]string
	//在使用map之前，需要先make，make的作用就是给map分配数据空间
	//map key-value是无序的
	a = make(map[string]string, 10)
	a["no1"] = "宋江"//如果没有no1就是增加
	a["no2"] = "无用"
	a["no1"] = "吴广源"//如果有no1就是修改
	a["no3"] = "出发点"
	fmt.Println(a)

	//第二种使用方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "深圳"
	cities["no3"] = "广州"
	fmt.Println(cities)

	//第三种使用方式
	var heroes map[string]string = map[string]string{
		"hero1" : "宋江",
		"hero2" : "吴广源",
	}
	fmt.Println(heroes)
	//heroes := map[string]string{
	//	"hero1" : "宋江",
	//	"hero2" : "吴广源",
	//}

	//案例
	/*
	课堂练习：演示一个key-value 的value是map的案例
	比如：我们要存放3个学生信息，每个学生有name，sex信息
	思路 ： map[string]map[string]string
	 */

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "赤沙南约新街"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "荷兰"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
	fmt.Println(studentMap["stu02"]["address"])

}
