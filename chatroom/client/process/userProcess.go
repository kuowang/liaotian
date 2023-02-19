package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"liaotian/chatroom/client/common/message"
	"liaotian/chatroom/server/utils"
	"net"
)

type UserProcess struct {
}

// 登录用户
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	fmt.Println("用户名密码", userId, userPwd)
	//链接服务器
	conn, err := net.Dial("tcp", "localhost:8880")
	if err != nil {
		fmt.Println("链接服务器失败", err)
		return err
	}
	defer conn.Close()

	//创建发送消息结构体
	var msg message.Message
	msg.Type = message.LoginMesType
	//登录消息内容结构体
	var loginMse message.LoginMes
	loginMse.UserId = userId
	loginMse.UserPwd = userPwd
	//将消息内容序列号
	data, err := json.Marshal(loginMse)
	if err != nil {
		fmt.Println("json 序列号失败", err)
	}
	msg.Data = string(data)
	//再次序列号消息体数据
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json 消息体序列化失败", err)
		return
	}

	//发送数据前需要整理数据
	// 1将data的长度发送给服务器
	// 2需要先获取data的长度 再转换成byte切片
	var pkglen uint32
	pkglen = uint32(len(data)) //转成无符号整形
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)

	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("数据发送异常", err)
	}
	fmt.Println("发送消息数据长度:", len(data), "发送内容:", string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("发送数据异常", err)
	}

	//time.Sleep(5 * time.Second)
	//fmt.Println("休息一会关闭 客户端")

	ut := utils.Transfer{Conn: conn}

	msg, err = ut.ReadPkg()
	if err != nil {
		fmt.Println("读取数据异常")
		return
	}
	//msg 反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(msg.Data), &loginResMes)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	if loginResMes.Code == 200 {
		fmt.Println("用户登录成功")
		//初始化curUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		fmt.Println("初始化链接")

		fmt.Println("当前在线用户列表")
		for _, v := range loginResMes.Users {
			fmt.Println(v)
			if v == userId {
				continue
			}
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUser[v] = user
		}
		//开启一个协程 处理服务器端给客户端的消息
		go ProcessServerMes(conn)
		for {
			ShowMenu()
		}

	} else {
		fmt.Println("用户登录失败", loginResMes.Error)
		return err
	}
	return nil
}

// 登录用户
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	fmt.Println("用户名密码", userId, userPwd, userName)
	//链接服务器
	conn, err := net.Dial("tcp", "localhost:8880")
	if err != nil {
		fmt.Println("链接服务器失败", err)
		return err
	}
	defer conn.Close()

	//创建发送消息结构体
	var msg message.Message
	msg.Type = message.RegisterMesType
	//登录消息内容结构体
	var RegisterMes message.RegisterMes
	RegisterMes.User.UserId = userId
	RegisterMes.User.UserPwd = userPwd
	RegisterMes.User.UserName = userName
	//将消息内容序列号
	data, err := json.Marshal(RegisterMes)
	if err != nil {
		fmt.Println("json 序列号失败", err)
	}
	msg.Data = string(data)
	//再次序列号消息体数据
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json 消息体序列化失败", err)
		return
	}

	//发送数据前需要整理数据
	// 1将data的长度发送给服务器
	// 2需要先获取data的长度 再转换成byte切片
	var pkglen uint32
	pkglen = uint32(len(data)) //转成无符号整形
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)

	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("数据发送异常", err)
	}
	fmt.Println("发送消息数据长度:", len(data), "发送内容:", string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("发送数据异常", err)
	}

	//time.Sleep(5 * time.Second)
	//fmt.Println("休息一会关闭 客户端")

	ut := utils.Transfer{Conn: conn}

	msg, err = ut.ReadPkg()
	if err != nil {
		fmt.Println("读取数据异常")
		return
	}
	//msg 反序列化
	var RegisterResMes message.RegisterResMes
	err = json.Unmarshal([]byte(msg.Data), &RegisterResMes)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	if RegisterResMes.Code == 200 {
		fmt.Println("用户注册成功")

		//开启一个协程 处理服务器端给客户端的消息
		go ProcessServerMes(conn)
		for {
			ShowMenu()
		}

	} else {
		fmt.Println("用户注册失败", RegisterResMes.Error)
		return err
	}
	return nil
}

// 保持和服务器端的消息通讯
func ProcessServerMes(conn net.Conn) {
	//J创建要给实例 读取服务器的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("客户端等待读取服务器发送消息")
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType: //用户上线通知
			var notifyUserMes = message.NotifyUserStatusMes{}
			json.Unmarshal([]byte(mes.Data), notifyUserMes)

			UpdataUserStatus(&notifyUserMes)
			outputOnlineUser()

		case message.SmsMesType: //在线消息内容展示
			outputGroupMes(&mes)

		default:
			fmt.Println("返回的内容无法识别", mes)
		}

		//fmt.Println(mes)
	}

}
