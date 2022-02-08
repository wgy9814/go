package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//通过go 向redis 写入数据和读取数据
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	// 关闭连接
	defer conn.Close()
	//2. 通过go向redis写入数据 string [key - value]
	_, err = conn.Do("Set", "name", "Tom")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//3. 通过go向redis读取数据 string [key - value]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	//因为返回 r 是interface{}
	//因为 name 对应的值是string，因此我们需要转换
	//nameString := r.(string)
	fmt.Println("操作ok ", r)

}