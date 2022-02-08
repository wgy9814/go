package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Monster struct {
	Name string
	Age int
	Skill string
}
func Process(name string, age int, skill string) {

	// 1.连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn err = ", err)
		return
	}
	defer conn.Close()

	// 2.往redis数据库写入数据
	_, err = conn.Do("HMset", "monster", "name", name, "age", age, "skill", skill)
	if err != nil {
		fmt.Println("write data err = ", err)
		return
	}

	// 3.从redis中读取数据
	r, err := redis.Strings(conn.Do("HMget", "monster", "name", "age", "skill"))
	if err != nil {
		fmt.Println("read data err = ", err)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%d]=%v\n", i, v)
	}

}

func main() {
	//Monster信息[name, age, skill]
	//通过终端输入三个monster的信息，使用golang操作redis，存放到redis中[用hash数据类型]
	//遍历出所有的Monster信息，并显示在终端
	//保存monster可以使用hash数据类型，遍历时先取出所有的keys，比如keys monster*

	// 定义结构体实例
	var monster Monster
	fmt.Println("请输入monster的Name：")
	fmt.Scanln(&monster.Name)
	fmt.Println("请输入monster的Age：")
	fmt.Scanln(&monster.Age)
	fmt.Println("请输入monster的Skill：")
	fmt.Scanln(&monster.Skill)

	Process(monster.Name, monster.Age, monster.Skill)

}