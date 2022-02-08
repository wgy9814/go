package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//打开一个存在的文件，将原来的内容覆盖成新的内容10句 "你好，尚硅谷"
	filePath := "g:/go/files/test/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=", err)
		return
	}


	//及时关闭file句柄
	defer file.Close()

	//准备写入5句“hello gardon”
	str := "你好，尚硅谷\r\n"
	//写入时，使用带缓存的 *Write
	writer := bufio.NewWriter(file)
	for i := 0; i < 10 ; i++ {
		writer.WriteString(str)
	}

	//因为 write是带缓存，因此在调用WriteString方法时，其实内容
	//是先写到缓存的，所以需要调用Flush方法，将缓冲的数据
	//真正写到文件中，否则文件中会没有数据

	writer.Flush()
}