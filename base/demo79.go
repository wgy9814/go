package main

import (
	"fmt"
	"go_code/chapter/demo04/model"
)

func main() {
	account := model.NewAccount("454", "123456", 300.0)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败=", account)
	}
}
