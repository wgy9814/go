package main

import "fmt"

// 写数据
func WriteData(intChan chan int){
	for i := 1; i <= 50; i++{
		intChan<-i
		fmt.Println("写入管道数据:", i)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool){
	for  {
		v, ok := <-intChan
		if !ok{
			break
		}
		// 频率不一样完全没问题
		//time.Sleep(time.Second)
		fmt.Println("读到管道数据:", v)
	}
	exitChan<- true
	close(exitChan)
}


func main() {
	//请完成goroutine和channel协同工作的案例，具体要求:
	//开启一个writeData协程，向管道intChan中写入50个整数
	//开启一个readData协程，从管道intChan中读取writeData写入的数据。
	//注意: writeData和readDate操作的是同一个管道
	//主线程需要等待writeData和readDate协程都完成工作才能退出[管道]

	//阻塞的问题的提出。
	//问题:如果注销掉go readData(intq han, exitChan) ,程序会怎么样?
	//如果只是向管道写入数据，而没有读取，就会出现阻塞而dead lock，原因是intChan容量是10,而代码writeData会写入50个数据因此会阻塞在writeData的 ch<- i
	//如果，编译器(运行),发现一个管道、只有写，而没有读，则该管道，会阻塞。
	//写管道和读管道的频率不一致，无所谓。(编译器会自己分析,有没有在读没有报错死锁，有正常运行)

	// 创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	// 启动 写读协程
	go WriteData(intChan)
	go readData(intChan, exitChan)
	// 不停读取退出管道, 解决之前不知道何时协程执行完成的问题
	for {
		// 直到可以读到exitChan的值
		_, ok := <-exitChan
		// fmt.Println(ok) 这里!ok和ok是一样的结果
		// 因为!ok指的取到后再取一次取不到, ok是直接取到(更直接)
		if !ok{
			break
		}
	}
}