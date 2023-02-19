package message

//定义用户结构体

type User struct {
	//确定用户表字段 保证表的字段格式化
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
}
