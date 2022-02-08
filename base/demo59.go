package main

import "fmt"

func main() {
	//演示map切片的使用
	/*
	要求：使用一个map来记录monster的信息 name和age，也就是说一个
	monster对应一个map，并且妖怪的的个数可以动态的增加=》map切片
	 */
	//1. 声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2) //放入2个妖怪
	//2.增加第一个妖怪信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "孙悟空"
		monsters[1]["age"] = "1000"
	}

	//下面这个写法越界
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string, 2)
	//	monsters[2]["name"] = "孙悟空"
	//	monsters[2]["age"] = "1000"
	//}

	//这里我们需要使用到切片的append函数，可以动态增加monsters
	//1.先定义个monsters信息
	newMonster := map[string]string{
		"name" : "武德星君",
		"age" : "600",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)
}
