package main

import (
	"fmt"
	"go_code/chapter/demo04/model"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(50)
	p.SetSal(12000.0)
	fmt.Println(p)
	fmt.Println(p.Name, " age=", p.GetAge(), " sal=", p.GetSal())
}
