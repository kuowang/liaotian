package main

import (
	"fmt"
	"liaotian/chatroom/server/model"
	"net"
	"time"
)

func init() {
	//初始化redis连接池
	initPool("127.0.0.1:6379", 16, 0, time.Second*300)
	initUserDao() //初始化用户配置对象
}
func main() {

	fmt.Println("服务器开启监听:8889")
	listen, err := net.Listen("tcp", "0.0.0.0:8880")
	if err != nil {
		fmt.Println("监听开启失败", err)
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端链接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("链接失败", err)
			continue
		}
		//链接成功 启动协程和客户端保持通讯
		go GoProcess(conn)
	}
}

func GoProcess(conn net.Conn) {
	defer conn.Close() //延迟关闭
	pro := &Processor{
		Conn: conn,
	}
	err := pro.GoProcess()
	if err != nil {
		fmt.Println("协程异常,退出", err)
		return
	}
}

// 初始化userDao的链接
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
