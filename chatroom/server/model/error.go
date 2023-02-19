package model

import "errors"

//定义错误信息

var (
	ERROR_USER_NOTEXIESTS = errors.New("用户不存在..")
	ERROR_USER_EXISTS     = errors.New("用户已存在")
	ERROR_USER_PWD_ERROR  = errors.New("密码错误")
)
