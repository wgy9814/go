package message


const (
	LoginMesType  = "LoginMes"
	LoginResMesType  =  "LoginResMes"
	RegisterMesType   =  "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType          = "SmsMes"
)

//这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	Useroffline
	UserBusyStatus
)

type Message struct {
	Type string  `json:"type"` // 消息类型
	Data string `json:"data"`  // 消息具体内容
}


// 定义两种消息

type LoginMes struct {
	UserId int  `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"user_name"`
}


type LoginResMes struct {
	Code int   `json:"code"`  //状态码 500表示未注册 200表示登录成功
	UserId []int			  //增加字段，保存用户id的切片
	Error string `json:"error"`  // 错误信息
}

type RegisterMes struct {
	User User `json:"user"` //类型就是 User 结构体
}


type RegisterResMes struct {
	Code int   `json:"code"`  //状态码 400表示该用户已经存在 200表示注册成功
	Error string `json:"error"`  // 错误信息
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"code"`    //用户id
	Status int `json:"status"`  // 用户的状态
}

//增加一个 SmsMes //发送的消息
type SmsMes struct {
	Content string `json:"content"` //内容
	User //匿名结构体 继承
}
