package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)
// 定义一个全局的pool
var pool *redis.Pool

// 创建连接池
func init()  {
	pool = &redis.Pool{
		MaxIdle: 8,  // 表示最大空闲连接数
		MaxActive: 0,	// 表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100,	// 最大空闲时间
		Dial: func() (redis.Conn, error) {  // 初始化连接的代码，连接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}


func main() {
	conn := pool.Get()  // 从连接池中取出一个连接
	defer conn.Close()  // 关闭连接池，一旦关闭连接池，就不能从连接池再取出连接
	_, err := conn.Do("set", "car", "BBB")

	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	r, err := redis.String(conn.Do("Get", "car"))
	if err != nil {
		fmt.Println("read Data err=", err)
		return
	}
	fmt.Println("r=", r)
	// 如果要从pool池中取出连接，一定要保证连接池是没有关闭的
	//下面代码会报错
	pool.Close()
	conn2 := pool.Get()  // 从连接池中取出一个连接
	_, err = conn2.Do("set", "car", "BBB")

	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	r, err = redis.String(conn2.Do("Get", "car"))
	if err != nil {
		fmt.Println("read Data err=", err)
		return
	}
	fmt.Println("r=", r)

}