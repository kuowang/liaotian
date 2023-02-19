package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 用户状态常量
const (
	UserOnline     = 0
	userOffline    = 1
	UserBusyStatus = 2
)

// Message 消息结构体
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息数据
}

// LoginMes 具体消息 登录消息
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //密码
	UserName string `json:"userName"` //用户名
}

// LoginResMes 登录返回信息
type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态码 500 用户未注册,200登录成功
	Users []int  `json:"users"` //当前在线用户的id切片
	Error string `json:"error"` //返回错误信息
}

// RegisterMes 注册消息 发送消息类型
type RegisterMes struct {
	User User `json:"user"`
}

// RegisterResMes 注册消息 接收消息类型
type RegisterResMes struct {
	Code  int    `json:"code"` //200成功 500错误
	Error string `json:"error"`
}

// 服务器发送推送消息类型
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户状态
}

type SmsMes struct {
	Content string `json:"content"` //发送的内容
	User    User   //发送人
}
