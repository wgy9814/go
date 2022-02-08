package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义一个结构体，用于存储统计结果
type CharCount struct {
	ChCount     int //记录英文个数
	NumCount    int //记录数字个数
	SpaceCount  int //记录空格个数
	OthersCount int //记录其他字符个数
}
func main() {

	/*
	思路：
	   1- 打开一个文件，创建一个reader
	   2- 每读取一行，就去统计该行有多少个字符，数字，空格等数量
	   3 - 然后将结果存到一个结构体内
	*/

	fileName := "g:/go/files/test/666.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("打开文件出错了... err =%v \n", err)
		return
	}

	defer file.Close()

	// 定义一个CharCount的实例
	var Count CharCount

	//创建一个带缓存的reader
	reader := bufio.NewReader(file)

	// 循环读取文件fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("文件内容如读取完成... message = %v \n", err)
			break
		}
		// 遍历前面循环读取到的内容str, 进行统计数量
		for _, v := range str {
			//fmt.Println(v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				Count.ChCount++
			case v == ' ' || v == '\t':
				Count.SpaceCount++
			case v >= '0' && v <= '9':
				Count.NumCount++
			default:
				Count.OthersCount++
			}
		}
	}

	// 输出统计结果来看看是否正确
	fmt.Printf("字符数量= %v  数字的个数 = %v 空格的数量是= %v 另外其它字符的个数是= %v",
		Count.ChCount, Count.NumCount, Count.SpaceCount, Count.OthersCount)
}