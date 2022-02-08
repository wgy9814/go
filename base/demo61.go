package main

import "fmt"

func modify(map1 map[int]int) {
	map1[1] = 900
}

//定义一个学生结构体
type Stu struct {
	Name string
	Age int
	Address string
}

func main() {
	//map是引用类型，遵守引用类型传递的机制，在函数接收一个map
	//修改后，会直接修改原来的map

	//map能自动扩容
	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	//map是引用类型
	fmt.Println(map1)

	//map的value 也经常用struct类型
	//更适合管理复杂的数据（比前面value是一个map更好）
	//比如value为Student结构体
	//1.map的key为 学生学号，是惟一的
	//2.map的 value 为结构体，包含学生的名字，年龄，地址
	students := make(map[string]Stu, 10)
	stud1 := Stu{"jack", 18, "广州市海珠区"}
	stud2 := Stu{"mary", 30, "美国佛罗里达"}
	students["no1"] = stud1
	students["no2"] = stud2

	fmt.Println(students)

	//遍历各个学生信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v \n", k)
		fmt.Printf("学生的名字是%v \n", v.Name)
		fmt.Printf("学生的年龄是%v \n", v.Age)
		fmt.Printf("学生的地址是%v \n", v.Address)
		fmt.Println()
	}


}
