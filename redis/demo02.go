package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 通过go向redis写入数据和读取数据
	// 1.连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis connect failed. err = ", err)
		return
	}

	defer conn.Close() //关闭连接

	// 2.通过go向redis写入数据
	_, err = conn.Do("HMSet", "user01", "name", "jack", "age", 20)
	if err != nil {
		fmt.Println("HMSet err = ", err)
		return
	}

	// 3.从redis中读取数据
	// 由于返回的r是一个多类型的变量，故需要用redis.Strings()
	r, err := redis.Strings(conn.Do("HMGet", "user01", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err = ", err)
		return
	}
	// 遍历结果
	for i, v := range r {
		fmt.Printf("r[%d]=%v\n", i, v)
	}
	/*
		r[0]=jack
		r[1]=20
	*/

}