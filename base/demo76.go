package main

import (
	"fmt"
	"go_code/chapter/demo04/model"
)


/*

 */
type Stu struct {
	Name string
	Age int
}



func main() {
	//创建要给 Student 实例
	//var stu = model.Student{
	//	Name : "tom",
	//	Score : 20,
	//}
	//fmt.Println(stu)

	//因为 student 结构体首字母是小写,通过工厂模式来解决
	var stu =model.NewStident("tom-", 88)
	fmt.Println(*stu)
	//fmt.Println("name=", stu.Name, "score=", stu.Score)

	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
