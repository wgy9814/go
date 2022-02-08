package main

import (
	"fmt"
	"reflect"
)

// 普通数据类型的反射
func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量的type(类型)、kind(类别)值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType:", rTyp)
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal:", rVal)
	// rVal直接输出虽然是100 但是实际上它的类型是reflect.Value类型 查看手册看函数
	fmt.Printf("rVal : %v, rVal Type : %T \n", rTyp, rTyp)
	n2 := 2 + rVal.Int()
	fmt.Println("n2:=", n2)
	// 3. 我们将rVal转为interface{}
	iV := rVal.Interface()
	// 将interface通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2:%v, num2的类型%T \n", num2, num2)


}

//对结构体的反射
func reflectTest02(b interface{}){
	//通过反射获取的传入的变量的type(类型)、kind(类别)值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType:", rTyp)
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal:", rVal)
	fmt.Printf("rVal : %v, rVal Type : %T \n", rTyp, rTyp)
	//2.1获取变量对应的kind
	//（1）rVal.Kind() ==>
	kind1 := rVal.Kind()
	//（2）rTyp.Kind() ==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind = %v, kind = %v \n", kind1, kind2)
	// 3. 我们将rVal转为interface{}
	iV := rVal.Interface()
	//将interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//fmt.Printf("iV:%v, iV的类型%T, iV Name的值%v \n", iV, iV, iV.Name)
	// 这里在运行时可以知道是Student类型 但是编译器并不知道它是哪个类型
	fmt.Printf("iV:%v, iV的类型%T \n", iV, iV)
	// 这里必须对它断言 获取它的值
	// 这里可以使用swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu:%v, stu的类型%T, stu Name的值%v \n", stu, stu, stu.Name)
	}
}

type Student struct{
	Name string
	Age int
}

func main() {
	//请编写一个案例，
	//演示对(基本数据类型、interface{}、 reflect.Value)进行反射的基本操作
	//1.先定义一个int
	var num int
	num = 100
	reflectTest01(num)

	//2.定义一个Students的实例
	stu := Student{
		Name: "Tom",
		Age: 20,
	}
	reflectTest02(stu)


}