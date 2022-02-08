package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 1.连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis connect failed. err = ", err)
		return
	}

	//// 设置数据有效时间
	//_, err = conn.Do("MSet", "name01", "RSQ", "name02", "Tom")
	//if err != nil {
	//	fmt.Println("HMSet err = ", err)
	//	return
	//}
	//// 设置超时时间
	//_, err = conn.Do("expire", "name01", 10)
	//if err != nil {
	//	fmt.Println("Set Expire err = ", err)
	//	return
	//}

	// 1.通过go向redis写入数据
	_, err = conn.Do("Lpush", "herolist", "no1:宋江", 30, "no2:吴淞", 40)
	if err != nil {
		fmt.Println("Lpush err = ", err)
		return
	}

	// 2.从redis中读取数据
	r, err := redis.String(conn.Do("rpop", "herolist"))
	if err != nil {
		fmt.Println("rpop err = ", err)
		return
	}


}