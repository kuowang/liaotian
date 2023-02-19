package process

import (
	"fmt"
	"liaotian/chatroom/client/common/message"
	"liaotian/chatroom/client/model"
)

// 维护客户端的map
var onlineUser map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

func UpdataUserStatus(mes *message.NotifyUserStatusMes) {

	user, ok := onlineUser[mes.UserId]
	if !ok {
		user = &message.User{
			UserId: mes.UserId,
		}
	} else {
		user = &message.User{
			UserId:     mes.UserId,
			UserStatus: mes.Status,
		}
	}

	onlineUser[mes.UserId] = user
}

func outputOnlineUser() {
	fmt.Println("当前用户列表")
	for id, user := range onlineUser {
		if id == userId {
			continue
		} else {
			fmt.Println("用户上线id:", id, user)
		}

	}
}
