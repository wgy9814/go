package main

import (
	"fmt"
	"strconv"
	"time"
)

//每隔1秒输出"hello,world"
func test(){
	for i:=0; i < 10; i++{
		fmt.Println("test() hello, world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

//在主线程(可以理解成进程)中，开启一个goroutine,该协程每隔1秒输出"hello,world"
//在主线程中也每隔一秒输出"hello,golang",输出10次后，退出程序
//要求主线程和goroutine同时执行
func main() {
	go test() //开启了一个协程
	for i:=0; i < 10; i++{
		fmt.Println("test() hello, Golang mian" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}