package process

//在线用户集合
import "fmt"

var (
	usreMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	usreMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// AddOnlineUser 添加map 在线用户
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// DelOnlineUser 删除map在线用户
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// GetAllOnlineUser 获取所有的在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// GetOnlineUserById 获取指定的用户
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户不在线")
		return
	}
	return
}
