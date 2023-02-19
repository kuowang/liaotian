package main

import (
	"errors"
	"fmt"
	"io"
	"liaotian/chatroom/client/common/message"
	"liaotian/chatroom/server/process"
	"liaotian/chatroom/server/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 根据消息体类型调用不同的函数
func (this *Processor) serverProcessMsg(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType: //处理登录
		userProcess := &process.UserProcess{Conn: this.Conn}
		err := userProcess.ServerProcessLogin(mes)

		if err != nil {
			return err
		}

	case message.RegisterMesType: //注册信息
		userProcess := &process.UserProcess{Conn: this.Conn}
		_, err := userProcess.ServerProcesRegister(mes)

		if err != nil {
			return err
		}
	case message.SmsMesType:
		smsProcess := &process.SmsProcess{}
		smsProcess.SendMes(mes)

		fmt.Println("接受群聊消息", mes)
	//
	default:
		fmt.Println("类型不存在,无法处理")
		err = errors.New("类型不存在,无法处理")
		return

	}

	return nil
}

func (this *Processor) GoProcess() (err error) {
	//接收数据
	for {
		//将读取的数据包 直接封装成函数
		tf := utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务器正常退出")
				return err
			}
			fmt.Println("读取数据异常", err)
		}
		fmt.Println(mes)

		err = this.serverProcessMsg(&mes)
		if err != nil {
			return err
		}
	}
	return err
}
