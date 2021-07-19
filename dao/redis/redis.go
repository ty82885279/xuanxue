package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"xuanxue/settings"
)

// 声明一个全局RedisDB
var Db *redis.Client

func Init(cfg *settings.RedisConf) (err error) {

	Db = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})
	_, err = Db.Ping().Result()
	return
}
