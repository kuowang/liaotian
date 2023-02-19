package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxidle int, maxActive int, idleTimeOut time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxidle,     //最大空闲连接数
		MaxActive:   maxActive,   //最大连接数
		IdleTimeout: idleTimeOut, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
	//c:=pool.Get()
}
