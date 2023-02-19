package main

import (
	"fmt"
	"liaotian/chatroom/client/process"
)

// 定义用户id和密码
var userId int
var userPwd string
var userName string

func main() {
	//接受用户选择
	var key int16
	//判断是否选择显示菜单
	var loop = true

	for {
		fmt.Println("----欢迎进入聊天室------")
		fmt.Println("\t 1 登录聊天室")
		fmt.Println("\t 2 注册用户")
		fmt.Println("\t 3 退出系统")
		fmt.Println("\t 请选择1-3")

		fmt.Println(key, loop)

		//获取输入数据
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1: //登录
			login()
		case 2: //注册
			//fmt.Println("注册用户")
			register()
		case 3:
			fmt.Println("退出系统1")
		default:
			fmt.Println("输入有误,请重新输入")
		}

	}

}

func login() {
	fmt.Println("登陆聊天室")
	fmt.Println("请输入用户id")
	fmt.Scanf("%d\n", &userId)
	fmt.Println("请输入密码")
	fmt.Scanf("%s\n", &userPwd)
	//完成登录
	up := &process.UserProcess{}
	err := up.Login(userId, userPwd)
	if err == nil {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}

}

func register() {
	fmt.Println("注册用户")
	fmt.Println("请输入用户id")
	fmt.Scanf("%d\n", &userId)
	fmt.Println("请输入用密码")
	fmt.Scanf("%s\n", &userPwd)
	fmt.Println("请输入用户名")
	fmt.Scanf("%s\n", &userName)

	//完成登录
	up := &process.UserProcess{}
	err := up.Register(userId, userPwd, userName)
	if err == nil {
		fmt.Println("登录成功")
	} else {
		fmt.Println("登录失败")
	}

}
