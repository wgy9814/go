package main

import (
	"encoding/json"
	"fmt"
)

//结构体
type A struct {
	Num int
}

//结构体
type B struct {
	Num int
	//Name string
}

type Monster struct {
	Name string `json:"name"` //`json:"name"` 就是 struct tag
	Age int `json:"age"`
	Skill string `json:"skill"`
}


type Student struct {
	Name string
	Age int
}


func main() {
	var a A
	var b B
	a = A(b)//可以转换，但是有要求，就是结构体的字段要完全一样（包括名字，个数和类型）
	fmt.Println(a,b)

	//结构体进行type重新定义（相当于取别名），go认为是新的数据类型，但是相互间可以强转
	type Stu Student

	var stu1 Student
	var stu2 Stu
	stu2 = Stu(stu1)
	fmt.Println(stu1,stu2)

	//1.创建一个变量
	monster := Monster{"牛魔王", 500, "芭蕉扇"}
	//2.将变量序列化为json格式字串
	//json.Marshal 函数使用反射
	jsonstr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误",err)
	}
	fmt.Println("jsonstr",string(jsonstr))

}
