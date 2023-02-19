package model

import (
	"liaotian/chatroom/client/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
