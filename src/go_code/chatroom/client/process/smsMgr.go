package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	//显示即可
	//1.反序列化mes.data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t %s",
		smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
