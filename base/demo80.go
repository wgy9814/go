package main

import (
	"fmt"
)

//编写一个学生考试系统
type Student struct {
	Name string
	Age int
	Score int
}

//将 Pupil 和 Graduate 共有的方法也绑定到 *Student

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	//业务判断
	stu.Score = score
}

//给 *Student 增加一个方法 那么 Pupil 和 Graduate都可以用这个方法
func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

//小学生
type Pupil struct {
	Student //嵌入了Student匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试....")
}

//小学生
type Graduate struct {
	Student //嵌入了Student匿名结构体
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试....")
}


func main() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom-"
	pupil.Student.Age = 7
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
	fmt.Println("res=", pupil.Student.GetSum(1, 2))

	graduate := &Graduate{}
	graduate.Student.Name = "jack-"
	graduate.Student.Age = 23
	graduate.testing()
	graduate.Student.SetScore(80)
	graduate.Student.ShowInfo()
	fmt.Println("res=", graduate.Student.GetSum(10, 20))
}
