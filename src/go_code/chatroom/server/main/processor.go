package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

//这里将这些方法关联到结构体
type Processor struct {
	//分析他应该有哪些字段
	Conn net.Conn
}


//根据客户端发送消息类型不同，决定调用哪个函数来处理
func (this *Processor) severProcessMes(mes *message.Message) (err error){
	//看看我们能不能收到客户端发送的群发消息
	fmt.Println("mes=", mes)
	switch mes.Type {
		case message.LoginMesType:
			//处理登录
			//创建一个 UserProcess 实例
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.SeverProcessLogin(mes)
		case message.RegisterMesType:
			//处理注册
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.SeverProcessRegister(mes)
		case message.SmsMesType:
			//创建一个SmsProcess实例完成群聊功能
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default:
			fmt.Println("消息类型不存在，无法处理....")
	}
	return
}

func (this *Processor) process2() (err error){
	// 循环读取客户端发送的信息
	for  {
		//这里我们将读取数据包 直接封装成一个函数 readPkg(),返回 Message， err
		//创建一个 Transfer 实例，然后读取
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出，服务器端也退出...")
				return err
			}
			//fmt.Println("readPkg err=", err)
			return err
		}
		//fmt.Println("客户端发送的消息为:",message)
		err = this.severProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}