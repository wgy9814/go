package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess) Register(userId int, userPwd string,
	userName string) (err error) {
	// 1.链接到服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("链接错误", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2. 通过conn,发送消息给服务端
	var mess message.Message
	mess.Type = message.RegisterMesType

	// 3.创建一个LoginMes 结构体,并赋值
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	// 4.将结构体 LoginMes序列化成字符串
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("序列化错误", err)
		return
	}
	// 5.将dara赋值给 mess.Data 将字节切片转成字符串
	mess.Data = string(data)

	//6. 将mess 结构体序列化成字符串
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("序列化错误", err)
		return
	}

	//创建一个 Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
		return
	}

	mes, err := tf.ReadPkg() //mes 就是 RegisterResMes
	if err != nil {
		fmt.Println("readPkg(conn) fail", err)
		return
	}

	//将mes的Data部分序列化成 RegisterResMes
	var RegisterResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &RegisterResMes)

	if RegisterResMes.Code == 200 {
		fmt.Println("注册成功，你可以重新登陆")
		os.Exit(0)
	} else {
		fmt.Println(RegisterResMes.Error)
		os.Exit(0)
	}

	//fmt.Printf("客户端。发送消息的长度=%d", len(data))
	//fmt.Println(n, string(data))
	return

}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//fmt.Printf("userId = %d password=%s\n", userId, password)
	//return nil
	// 1.链接到服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("链接错误", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2. 通过conn,发送消息给服务端
	var mess message.Message
	mess.Type = message.LoginMesType

	// 3.创建一个LoginMes 结构体,并赋值
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//loginMes.UserName = "xiaoming"

	// 4.将结构体 LoginMes序列化成字符串
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化错误", err)
		return
	}
	// 5.将dara赋值给 mess.Data 将字节切片转成字符串
	mess.Data = string(data)

	//6. 将mess 结构体序列化成字符串
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("序列化错误", err)
		return
	}

	//7. 到这个时候 data就是我们要发送的消息
	// 先发送data 的字节长度,发送给 服务器,防止丢包
	// 先获取data 长度,转换 成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buffer [4]byte

	binary.BigEndian.PutUint32(buffer[0:4], pkgLen)

	n, err := conn.Write(buffer[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Printf("客户端，发送消息的长度=%d 内容=%s", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write data fail", err)
		return
	}

	//休眠10秒
	//time.Sleep(10 * time.Second)
	//fmt.Println("休眠了10秒")

	//处理服务器端返回的消息
	//创建一个 Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err := tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) fail", err)
		return
	}

	//将mes的Data部分序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		//初始化 curUser
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline
		//fmt.Println("登录成功")

		//可以显示当前在线用户列表，遍历loginResMes.UserId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserId{
			//如果我们要求不显示自己在线，增加一个判断
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成客户端 的 onlineUsers 完成初始化
			user := &message.User{
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go severProcessMes(conn)

		//1.显示我们登录成功的菜单【循环】
		for  {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	//fmt.Printf("客户端。发送消息的长度=%d", len(data))
	//fmt.Println(n, string(data))
	return

}