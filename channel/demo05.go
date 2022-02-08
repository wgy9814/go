package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int){
	for i := 1; i <= 200000; i++{
		intChan<- i
	}
	close(intChan)
}

// 取值判断是否为素数
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool){
	var flag bool
	for  {
		//time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true //假设是素数
		// 判断是不是素数.
		for i := 2; i <num; i++ {
			if num % i == 0 {//说明该num不是素数
				flag = false
				break
			}
		}
		if flag{
			//将这个数放入到 primeChan
			primeChan<- num
		}
	}
	fmt.Printf("有协程已经取不到数据啦, 退出")
	exitChan<- true
}
func main() {
	//要求统计1-200000的数字中，哪些是素数? 学习了goroutine和channel的知识后，就可以完成了[测试数据: 8000]
	//传统的方法，就是使用一个循环，循环的判断各个数是不是素数[ok] 。
	//使用并发/并行的方式，将统计素数的任务分配给多个(4个)goroutine去完成，完成任务时间短。
	//协程后，执行的速度，比普通方法提高至少4倍

	intChan := make(chan int, 20000)
	primeChan := make(chan int, 10000)
	exitChan := make(chan bool, 4)

	// 统计用时
	start := time.Now().Unix()
	// 开启一个协程放入1-8000个数
	go putNum(intChan)
	// 开启四个协程取出数据判断素数
	for i := 0; i < 4; i++{
		go primeNum(intChan, primeChan, exitChan)
	}

	// 再开一个协程 等待标志管道取出四个值
	go func(){
		for i := 0; i < 4; i++{
			<- exitChan
		}
		end := time.Now().Unix()
		fmt.Printf("\n使用协程耗时:%v", end - start)
		// 当我们从exitChan取出4个结果，就可以放心关闭管道 primeChan
		close(primeChan)
	}()
	// 遍历我们的primeChan把结果取出
	for  {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数%d\n", res)
	}
	fmt.Println("main线程退出")


}