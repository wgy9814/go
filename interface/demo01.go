package main

import "fmt"

//定义一个接口
type Usb interface {

	//声明了两个没有实现的方法
	Start()
	Stop()
}
type Usb2 interface {

	//声明了两个没有实现的方法
	Start()
	Stop()
	Test()
}

type Phone struct {
	Name string
	Score float64
}

//让 Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作.....")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作.....")
}

type Camera struct {

}

//让 Camera 实现 Usb接口的方法
func (p Camera) Start() {
	fmt.Println("相机开始工作.....")
}

func (p Camera) Stop() {
	fmt.Println("相机停止工作.....")
}

//计算机
type Computer struct {

}

//编写一个方法 Working 方法，接收一个Usb接口类型变量
//只要是实现了 Usb接口 （所谓的Usb接口，就是指实现了 Usb接口声明所有方法）

//usb 接口变量就体现出多态的特点
func (c Computer) Working(usb Usb) {
	//通过usb接口变量来调用Start 和 Stop 方法
	usb.Start()
	usb.Stop()
}

func main() {
	//测试
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}
