package models

import "github.com/garyburd/redigo/redis"

var (
	RedisPool *redis.Pool
)

func InitRedisPool() {
	RedisPool = &redis.Pool{
		MaxIdle:     16,  //最初的连接数量
		MaxActive:   0,   //最大的连接数量，0表示按照需要创建
		IdleTimeout: 300, //连接关闭时间
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func SetKey(key string, value interface{}) {
	conn := RedisPool.Get()
	defer conn.Close()
	conn.Do("SET", key, value)
}

func GetKey(key string) string {
	return ""
}
