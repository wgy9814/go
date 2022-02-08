package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	Name string `json:"monster_name"` //反射机制
	Age int `json:"monster_age"`
	Birthday string
	Sal float64
	Skill string
}

func testStruct() {
	//演示
	monster := Monster{
		Name : "牛魔王",
		Age : 500,
		Birthday : "10-20",
		Sal : 8000.0,
		Skill : "抱妹杀",
	}
	//将 monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("monster序列号后=%v\n", string(data))

}

//将map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	//使用map。需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "广州"

	//将 a这个map 序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("a序列化后=%v\n", string(data))

}

//将切片进行序列化
func testSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "张无忌"
	m1["age"] = 25
	m1["address"] = "冰火岛"
	slice = append(slice,m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "张三丰"
	m2["age"] = 89
	m2["address"] = [2]string{"武当山","夏威夷"}
	slice = append(slice,m2)

	//将切片进行序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	//输出序列化之后的结果
	fmt.Printf("序列化后=%v\n",string(data))
}

//将基本数据类型进行序列化（对基本数据类型进行序列化意义不大）
func testFloat64()  {
	var num float64 = 3456.85

	//对num进行序列化
	data, err := json.Marshal(num)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	//输出序列化之后的结果
	fmt.Printf("序列化后=%v\n",string(data))
}


func main() {
	//演示将结构体，map，切片进行序列号

	testStruct()
	testMap()
	testSlice()
	testFloat64()

}