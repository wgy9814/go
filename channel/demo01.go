package main

import "fmt"

func main() {
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	// 2. 向管道中写入数据
	intChan <- 10
	num := 30
	intChan <- num
	intChan <- 40
	// intChan <- 50 //不能超过最大长度 会报错all goroutines are asleep - deadlock!

	//3.看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) //3 3

	//4. 从管道中读取数据
	var num1 int
	num1 = <-intChan
	fmt.Println("num1:", num1)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) //3 3

	//5.在没有使用协程的情况下(取完没放入)，如果channel数据取完了，再取，就会报deadlock
	num2 := <-intChan
	num3 := <-intChan
	//如果从channel取出数据后，可以继续放入
	//num4 := <-intChan //如果管道中无值再取 报错all goroutines are asleep - deadlock!
	//fmt.Println("num4:", num4)
	fmt.Println("num1:", num1, "num2:", num2,"num3:", num3)

}