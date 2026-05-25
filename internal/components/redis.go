package components

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Redis *redis.Client

func InitRedis() {
	addr := viper.GetString("redis.addr")
	if addr == "" {
		panic("please add redis.addr in config")
	}
	pass := viper.GetString("redis.password")
	if pass == "" {
		panic("please add redis.password in config")
	}
	db := viper.GetString("redis.db")
	if db == "" {
		panic("please add redis.db in config")
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       cast.ToInt(db),
		PoolSize: 10,
	})
	zap.L().Info("redis addr: ", zap.String("addr", addr))
}
