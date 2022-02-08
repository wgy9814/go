package main

import (
	"fmt"
	"sync"
	"time"
)

//需求:现在要计算1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
//最后显示出来。要求使用goroutine完成。

//思路
//1.编写一个函数，来计算各个数的阶乘，并放入map中
//2.我们启动的协程多个，统计的结果放入到map中
//3.map 应该做出一个全局的

// 定义一个全局map
var(
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	//sync 是包：synchornized 同步
	//Mutex:是互斥
	lock sync.Mutex
)

// 计算n!结果以及之前结果保存在myMap中
func test(n int){
	res := 1
	for i := 1; i <= n; i++{
		res *= i
	}
	// 将结果储存在myMap中
	// 加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main(){
	//下面代码存在上面两个问题：
	//资源竞争
	//协程运行的退出(这里用Sleep延时肯定不行的)

	//我们这里开启多个协程完成这个任务【200个】
	for i := 1; i <= 20; i++{
		go test(i)
	}

	// 休眠10秒 防止主线程退出 其他协程也退出
	time.Sleep(time.Second*10)
	// 输出变量结果
	// 主线程并不知道10秒能执行完成，因此底层可能仍然出现资源争夺，因此加入互斥锁即可解决问题
	lock.Lock()
	for i, v := range myMap{
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()


}
