package process

import (
	"encoding/json"
	"fmt"
	"liaotian/chatroom/client/common/message"
	"liaotian/chatroom/server/model"
	"liaotian/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

// 处理登录请求信息
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//从mes中取出data 并反序列化loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("反序列化失败1", err)
	}
	//判断用户是否合法
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//再声明要给返回数据类型
	var loginResMes message.LoginResMes

	//去数据库验证用户
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		loginResMes.Code = 500
		if err == model.ERROR_USER_PWD_ERROR {
			loginResMes.Error = err.Error()
		}
		loginResMes.Error = err.Error()
	} else {
		loginResMes.Code = 200
		loginResMes.Error = ""
		this.UserId = loginMes.UserId            //记录当前用户的id
		usreMgr.AddOnlineUser(this)              //记录当前用户在线
		this.NotifyOthersOnlineUser(this.UserId) //通知其他用户自己上线

		for id, _ := range usreMgr.onlineUsers {
			loginResMes.Users = append(loginResMes.Users, id)
		}

		fmt.Println("登录成功", user)
	}

	/*if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
		loginResMes.Error = ""
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户账号密码不正确"
	}*/
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	resMes.Data = string(data)
	//再对消息体序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	//发送数据,封装到函数中
	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	return
}

// 注册用户
func (this UserProcess) ServerProcesRegister(mes *message.Message) (user *model.User, err error) {

	//从mes中取出data 并反序列化loginMes
	var RegisterMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &RegisterMes)
	if err != nil {
		fmt.Println("反序列化失败1", err)
	}
	//判断用户是否合法
	var resMes message.Message
	resMes.Type = message.RegisterMesType
	//再声明要给返回数据类型
	var RegisterResMes message.RegisterResMes

	//去数据库验证用户
	user, err = model.MyUserDao.Register(&RegisterMes.User)
	if err != nil {
		RegisterResMes.Code = 500
		if err == model.ERROR_USER_PWD_ERROR {
			RegisterResMes.Error = err.Error()
		}
		RegisterResMes.Error = err.Error()
	} else {
		RegisterResMes.Code = 200
		RegisterResMes.Error = ""
		fmt.Println("注册成功", user)
	}

	data, err := json.Marshal(RegisterResMes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	resMes.Data = string(data)
	//再对消息体序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	//发送数据,封装到函数中
	tf := &utils.Transfer{Conn: this.Conn}
	err = tf.WritePkg(data)
	return

}

func (this *UserProcess) NotifyOthersOnlineUser(userId int) {

	//循环发送给在线用户
	for id, up := range usreMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}

}

func (this *UserProcess) NotifyMeOnline(userId int) {

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.Status = message.UserOnline
	notifyUserStatusMes.UserId = userId

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	mes.Data = string(data)
	str, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败1", err)
		return
	}
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(str)
	if err != nil {
		fmt.Println("发送失败", err)
		return
	}
}
