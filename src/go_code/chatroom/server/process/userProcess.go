package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"go_code/chatroom/server/utils"
	"net"
)
//这里将这些方法关联到结构体
type UserProcess struct {
	//分析他应该有哪些字段
	Conn net.Conn
	//增加一个字段 表示改conn是哪个用户
	UserId int
}
//这里我们编写通知所有在线用户的方法
//userId 要通知其他的在线用户，我上线
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历 onlineUsers,然后一个一个的发送 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == userId {
			continue
		}
		//开始通知，单独写一个方法
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装我们的 NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将 notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Printf("json Marshal err=%v\n", err)
		return
	}

	//将序列化后的 notifyUserStatusMes 赋值给 mes.Data
	mes.Data = string(data)
	//5 .对mes 再次序列化 准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Printf("json Marshal err=%v\n", err)
		return
	}

	//创建一个 Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
	return
}

func (this *UserProcess) SeverProcessRegister(mes *message.Message) (err error){
	//1. 先从mes 中取出 mes.Data,并直接反序列化为 RegisterMes
	var RegisterMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &RegisterMes)
	if err != nil {
		fmt.Printf("json.unmarshal err=%v\n", err)
		return
	}

	//1.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//2.再声明一个 RegisterResMes
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库去完成注册
	//1.使用model.MyUserDao 去redis验证
	err = model.MyUserDao.Register(&RegisterMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505 //500 表示该用户不存在
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506 //500 表示该用户不存在
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200 //200 表示注册成功
	}

	//3.将 LoginResMes 序列号
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Printf("json Marshal err=%v\n", err)
		return
	}

	//4.将data赋值给 resMes
	resMes.Data = string(data)
	//5 .对resMes 进行序列号 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Printf("json Marshal err=%v\n", err)
		return
	}
	//6.发送data，我们将其封装到writePkg 函数
	//因为使用分层模式（mvc），我们先创建一个 Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
//专门处理登录请求
func (this *UserProcess) SeverProcessLogin(mes *message.Message) (err error){
	//1. 先从mes 中取出 mes.Data,并直接反序列化为 LoginMes
	var LoginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &LoginMes)
	if err != nil {
		fmt.Printf("json.unmarshal err=%v\n", err)
		return
	}

	//1.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个 LoginResMes
	var LoginResMes message.LoginResMes

	//我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao 去redis验证
	user, err := model.MyUserDao.Login(LoginMes.UserId, LoginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			LoginResMes.Code = 500 //500 表示该用户不存在
			LoginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			LoginResMes.Code = 300
			LoginResMes.Error = err.Error()
		} else {
			LoginResMes.Code = 505
			LoginResMes.Error = "服务器内部错误..."
		}

	} else {
		LoginResMes.Code = 200
		fmt.Println(user, "登录成功")
		//这里，因为用户登录成功，我们就把登录成功的放入到userMgr中
		//将登录成功的用户的userId 赋给 this
		this.UserId = LoginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他的在线用户 我上线了
		this.NotifyOthersOnlineUser(LoginMes.UserId)
		//将当前在线用户的id 放入到 LoginResMes.UserId
		//遍历 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			LoginResMes.UserId = append(LoginResMes.UserId, id)
		}
	}
	//如果用户id = 100， 密码=123456 ，认为合法，否则不合法
	//if LoginMes.UserId == 100 && LoginMes.Password == "123456" {
	//	//合法
	//	LoginResMes.Code = 200
	//} else {
	//	//不合法
	//	LoginResMes.Code = 500 //500 表示该用户不存在
	//	LoginResMes.Error = "该用户不存在，请注册再使用..."
	//}

	//3.将 LoginResMes 序列号
	data, err := json.Marshal(LoginResMes)
	if err != nil {
		fmt.Printf("json Marshal err=%v\n", err)
		return
	}

	//4.将data赋值给 resMes
	resMes.Data = string(data)
	//5 .对resMes 进行序列号 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Printf("json Marshal err=%v\n", err)
		return
	}
	//6.发送data，我们将其封装到writePkg 函数
	//因为使用分层模式（mvc），我们先创建一个 Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}