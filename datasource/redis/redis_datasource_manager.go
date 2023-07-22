package redis

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/switfssystem/switfs-framework/global"
)

type Config struct {
	Addr     string
	Port     string
	Password string
	DB       int
}

func New(cfg Config) {
	global.RedisClient = openConn(cfg)
	return
}

func openConn(redisCfg Config) *redis.Client {
	redisclient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,
	})
	return redisclient
}
