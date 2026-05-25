package components

import (
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Redis *redis.Client

func InitRedis() error {
	addr := viper.GetString("redis.addr")
	if addr == "" {
		return errors.New("missing redis.addr")
	}
	pass := viper.GetString("redis.password")
	db := viper.GetString("redis.db")
	if db == "" {
		return errors.New("missing redis.db")
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       cast.ToInt(db),
		PoolSize: 10,
	})
	zap.L().Info("redis configured", zap.String("addr", addr), zap.Int("db", cast.ToInt(db)))
	return nil
}
