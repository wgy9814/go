package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go_code/chatroom/common/message"
)

//我们在服务器启动后，就初始化一个userDao实例
//把他做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个 UserDao的结构体
//完成对User 结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个 UserDao 实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao){
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//思考一下在 UserDao 应该提供哪些方法给我们
//1. 根据用户id 返回一个 User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error){

	//提供给定id 去 redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误！
		if err == redis.ErrNil { //表示在 users 哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}

	//这里我们需要把 res 反序列化成 User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Printf("json.unmarshal err=%v\n", err)
	}

	return
}

//完成登录的校验 Login
//1.Login 完成对用户的验证
//2. 如果用户的id和pwd都正确。这返回一个 user 实例
//3. 如果用户的id和pwd错误。则返回对应的错误信息
func (this *UserDao) Login (userId int, userPwd string) (user *User, err error){
	//先从 UserDao 的连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//这是证明这个用户是获取到的
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

//完成登录的校验 Login
//1.Login 完成对用户的验证
//2. 如果用户的id和pwd都正确。这返回一个 user 实例
//3. 如果用户的id和pwd错误。则返回对应的错误信息
func (this *UserDao) Register (user *message.User) (err error){
	//先从 UserDao 的连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}


	//这时 说明id在redis还没有，则可以完成注册
	data, err := json.Marshal(user)//序列化
	if err != nil {
		return
	}

	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err =", err)
		return
	}
	return
}