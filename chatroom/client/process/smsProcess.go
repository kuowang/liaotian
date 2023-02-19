package process

import (
	"encoding/json"
	"fmt"
	"liaotian/chatroom/client/common/message"
	"liaotian/chatroom/server/utils"
)

type SmsProcess struct {
}

// 发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//创建message
	var mes = message.Message{}
	mes.Type = message.SmsMesType
	//创建发送内容实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.User = CurUser.User

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("序列化失败3", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败3", err)
		return
	}

	tf := &utils.Transfer{Conn: CurUser.Conn}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送失败2", err)
		return
	}
	return
}
