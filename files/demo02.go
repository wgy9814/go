package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//打开文件
	//概念说明：file的叫法
	//1. file 叫 file对象
	//2. file 叫 file指针
	//3. file 叫 file文件句柄
	file, err :=os.Open("g:/go/files/robots.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数关闭时，要及时的关闭fiel
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏

	//创建一个 *Reader, 是带缓冲的
	/*
	const (
		defaultBufsize = 4096 //默认的缓冲区为4096
	)
	 */
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读取到一个就换行
		if err == io.EOF { //io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}
	fmt.Println("文件读取结束....")

}