package process

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

//显示登录成功后的界面

func ShowMenu() {
	/* 显示在线用户列表

	 */
	fmt.Println("恭喜登录成功")
	fmt.Println("1显示在线用户列表")
	fmt.Println("2发送消息")
	fmt.Println("3信息列表")
	fmt.Println("4退出系统")

	fmt.Println("请选择1-4")

	var key int
	var content string
	s := &SmsProcess{} //发送消息实例

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入内容")
		fmt.Scanf("%s\n", &content)
		err := s.SendGroupMes(content)
		if err != nil {
			fmt.Println("发送失败")
			return
		}
	case 3:
	case 4:
		os.Exit(1)
	default:

	}

}
