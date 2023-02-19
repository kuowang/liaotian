package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"liaotian/chatroom/client/common/message"
)

var MyUserDao *UserDdao

type UserDdao struct {
	Pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDdao) {

	userDao = &UserDdao{
		Pool: pool,
	}
	return userDao
}

// 获取用户信息
func (this *UserDdao) GetUserInfo(conn redis.Conn, id int) (user *User, err error) {
	//通过id去redis中查询数据
	res, err := redis.String(conn.Do("HGet", "users", id))

	if err != nil {
		if err == redis.ErrNil {
			//没有查询到用户
			err = ERROR_USER_NOTEXIESTS
			return
		}
		return
	}
	fmt.Println("查询数据", res, err)
	user = &User{}
	err = json.Unmarshal([]byte(res), &user)

	if err != nil {
		fmt.Println("反序列化失败2")
		return
	}

	return
}

func (this *UserDdao) Login(userId int, userPwd string) (user *User, err error) {

	conn := this.Pool.Get()
	defer conn.Close()
	user, err = this.GetUserInfo(conn, userId)

	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD_ERROR
		return
	}
	return

}

func (this *UserDdao) Register(user *message.User) (userInfo *User, err error) {

	conn := this.Pool.Get()
	defer conn.Close()
	_, err = this.GetUserInfo(conn, user.UserId)

	if err == nil {
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("Hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保持注册用户错误", err)
		return
	}
	return

}
