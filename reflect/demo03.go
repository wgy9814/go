package main

import (
	"fmt"
	"reflect"
)


func main() {

	var str string = "tom"
	fs := reflect.ValueOf(&str)
	// fs的值:tom, fs的类型:reflect.Value
	fmt.Printf("fs的值:%v, fs的类型:%T\n", str, fs)
	// 这里fs需要是string或者v.CanSet()返回假
	fs.Elem().SetString("jack")
	fmt.Printf("fs的值:%v, fs的类型:%T\n", str, fs)
}