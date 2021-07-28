package model

import (
	"gift-code-Two/internal/globalError"
	"github.com/go-redis/redis"
)

//Rc redis操作
var Rc *redis.Client

//InitRc redis连接，连接之后会ping一下，如果ping失败，返回错误
func InitRc() error {
	Rc = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
		PoolSize: 50,
	})
	_, err := Rc.Ping().Result()
	if err != nil {
		return globalError.DBError("redis连接异常",globalError.RedisException)
	}
	return nil
}
