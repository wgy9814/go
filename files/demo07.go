package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//如果返回的错误为nil,说明文件或文件夹存在
//如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
//如果返回的错误为其它类型,则不确定是否在存在

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {

	//将 abc.txt 文件内容写入到 kkk.txt
	file1Path := "g:/go/files/test/abc.txt"
	file2Path := "g:/go/files/test/kkk.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("read file err=%v", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("write file err=%v", err)
		return
	}
}