package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 1; i <= 10; i++{
		time.Sleep(time.Second)
		fmt.Println("hello world sayHello", i)
	}
}

func test(){
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test拋出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	var myMap map[int]string
	// 没有make 就是用 肯定报错 看报错会不会终止主线程
	myMap[0] = "golang" //error
}

func main() {
	go sayHello()
	go test()
	for i := 0; i <= 10; i++{
		fmt.Println("main ok", i)
		time.Sleep(time.Second)
	}
}