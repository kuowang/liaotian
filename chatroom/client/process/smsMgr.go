package process

import (
	"encoding/json"
	"fmt"
	"liaotian/chatroom/client/common/message"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化失败，", err)
		return
	}
	//显示信息
	fmt.Println("用户ID：", smsMes.User.UserId)
	fmt.Println("消息：", smsMes.Content)

}
