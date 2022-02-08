package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
)

type SmsProcess struct {
	//[暂时不需要字段]
}

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端 的 onlineUsers map[int]*UserProcess
	//将消息转发出去

	//去除mes的内容 SmsMes
	var SmsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &SmsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers{
		//这里需要过滤掉自己，不要发给自己
		if id == SmsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个 Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
		return
	}
}