package utils

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"time"
)

var (
	RedisPool *redis.Pool
)

func InitRedisPool(){
	RedisPool = newRedisPool()
}

func newRedisPool() *redis.Pool {
	capacity := 10
	idleTimout :=  240 * time.Second
	network := "tcp"
	server := beego.AppConfig.String("redisServer")
	db := beego.AppConfig.String("redisDb")
	password := beego.AppConfig.String("redisPassword")
	return &redis.Pool{
		MaxIdle:     capacity,
		IdleTimeout: idleTimout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(network, server)
			if err != nil {
				return nil, err
			}

			if password != "" {
				_, err := conn.Do("AUTH", password)
				if err != nil {
					conn.Close()
					return nil, err
				}
			}

			if db != "" {
				_, err := conn.Do("SELECT", db)
				if err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}


func CloseRedisPool(){
	RedisPool.Close()
}
