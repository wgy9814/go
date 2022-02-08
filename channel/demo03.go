package main

import "fmt"

func main() {
	intchan := make(chan int, 3)
	intchan <- 10
	intchan <- 20
	// 关闭管道
	close(intchan)
	// 关闭后不能在存放
	//intchan <- 30
	// 读取数据完全没有问题
	n1 := <-intchan
	<- intchan
	fmt.Println(n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2<- i * 2 //放100个数据到管道
	}
	//	遍历管道不能使用普通的for循环 因为一个取出后第二个就变成第一个位置啦
	// for i := 0; i < len(intChan2); i++ {
	// }
	//在遍历时，如果channel没有关闭，则回出现deadlock的错误
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历。
	// 遍历前需要先关闭管道
	close(intChan2)
	for v :=range intChan2 {
		fmt.Println("v=", v)
	}

}