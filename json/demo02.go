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

//将json字符串，反序列化成struct
func unmarshalStruct() {
	//说明：str在项目的开发中，是通过网络传输获取到，或者读取文件获取到
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":500,\"Birthday\":\"10-20\",\"Sal\":8000,\"Skill\":\"抱妹杀\"}"

	//定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v\n", monster, monster.Name)
}

//将json字符串，反序列化成map
func unmarshalMap() {
	str := "{\"address\":\"广州\",\"age\":30,\"name\":\"红孩儿\"}"

	//定义一个map
	var a map[string]interface{}

	//注意：反序列化map,不需要make，因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)
}

//将json字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"冰火岛\",\"age\":25,\"name\":\"张无忌\"}," +
		"{\"address\":[\"武当山\",\"夏威夷\"],\"age\":89,\"name\":\"张三丰\"}]"


	//定义一个slice
	var slice []map[string]interface{}
	//注意：反序列化map,不需要make，因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}




func main() {
	//演示将json反序列化是指，将json字符串 反序列化成对应数据（比如：结构体，map，切片）

	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()


}