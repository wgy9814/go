package main

import (
	"fmt"
	"reflect"
	"testing"
)

//go语言通过反射创建结构体、赋值、并调用对应方法
type Call struct {
	Num1 int
	Num2 int
}

func (call Call) GetSub(name string){
	fmt.Printf("%v 完成了减法运算，%v - %v = %v \n", name, call.Num1, call.Num2, call.Num1 - call.Num2)
}

func (call *Call) GetSum(name string){
	fmt.Printf("%v 完成了加法运算，%v + %v = %v \n", name, call.Num1, call.Num2, call.Num1 + call.Num2)
}

func TestReflect(t *testing.T) {
	var (
		call *Call
		rValues []reflect.Value
		rValues2 []reflect.Value
	)
	ptrType := reflect.TypeOf(call) //获取call的指针的reflect.Type

	trueType := ptrType.Elem() //获取type的真实类型

	ptrValue := reflect.New(trueType) //返回对象的指针对应的reflect.Value

	call = ptrValue.Interface().(*Call)

	trueValue := ptrValue.Elem() //获取真实的结构体类型

	trueValue.FieldByName("Num1").SetInt(123)//设置对象属性，注意这个一定要是真实的结构类型的reflect.Value才能调用,指针类型reflect.Value的会报错
	//ptrValue.FieldByName("Num2").SetInt(23)
	trueValue.FieldByName("Num2").SetInt(23)

	//rValues = make([]reflect.Value, 0)
	rValues = append(rValues, reflect.ValueOf("xiaopeng"))//调用对应的方法
	fmt.Println(rValues)
	trueValue.MethodByName("GetSub").Call(rValues)
	/*
		fixme 在反射中，指针的方法不可以给实际类型调用，实际类型的方法可以给指针类型调用，因为go语言对这种操作做了封装
		所以下面一句是没问题的
		下下一句会运行时报错
	*/
	//ptrValue.MethodByName("GetSub").Call(rValues)
	//trueValue.MethodByName("GetSum").Call(append(rValues2, reflect.ValueOf("hiram")))
	ptrValue.MethodByName("GetSum").Call(append(rValues2, reflect.ValueOf("hiram")))

	fmt.Println(call)

	/*
		fixme 在实际使用中  指针和实体都能相互转换，不会影响调用
		但是指针的方法在方法体内的操作会影响到结构体本身属性
		而实体的方法不会，因为go对于结构体、数组、基本类型都是值传递
	*/
	call.GetSub("aaa")
	(*call).GetSub("bbb")
	call.GetSum("ccc")
	(*call).GetSum("ddd")
}