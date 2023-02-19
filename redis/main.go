package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //最大连接数
		IdleTimeout: 300, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	//c:=pool.Get()
}

func main() {
	conn := pool.Get()
	defer conn.Close()

}

func RedisConn() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		fmt.Println("链接失败")
		return
	}
	st, _ := redis.String(conn.Do("get", "user"))
	fmt.Println(st)
}
