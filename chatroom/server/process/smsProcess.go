package process

import (
	"encoding/json"
	"fmt"
	"liaotian/chatroom/client/common/message"
	"liaotian/chatroom/server/utils"
	"net"
)

type SmsProcess struct {
}

// 转发消息
func (this *SmsProcess) SendMes(mes *message.Message) {
	var SmsMes = message.SmsMes{}
	err := json.Unmarshal([]byte(mes.Data), &SmsMes)
	if err != nil {
		fmt.Println("反序列化失败4")
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败5")
	}

	for id, up := range usreMgr.onlineUsers {
		if id == SmsMes.User.UserId {
			continue
		}
		this.SendEachOnlineUser(data, up.Conn)
	}

}
func (this *SmsProcess) SendEachOnlineUser(data []byte, conn net.Conn) {

	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败", err)
	}

}
