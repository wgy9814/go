package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
func main() {
	//在执行test003之前 先获取当前unix时间戳
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("执行test003()好费时间=%v秒 \n", end-start)

}
