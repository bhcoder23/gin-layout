package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bhcoder23/gin-layout/internal/components"
	"github.com/bhcoder23/gin-layout/internal/interfaces"
	"github.com/bhcoder23/gin-layout/internal/models"
	"github.com/bhcoder23/gin-layout/internal/repos"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var Goods *GoodsService

type GoodsService struct {
	redis *redis.Client
	interfaces.Service[*models.Goods]
}

func NewGoodsService(r *repos.GoodsRepo, redisClient *redis.Client) *GoodsService {
	return &GoodsService{
		redis:   components.Redis,
		Service: *interfaces.NewService(repos.Goods),
	}
}

func init() {
	RegisterServices(func() {
		Goods = NewGoodsService(repos.Goods, components.Redis)
	})
}
func (s *GoodsService) GetByID(c *gin.Context) (goods *models.Goods, err error) {
	uid := c.GetUint(`goods_id`)
	key := fmt.Sprintf("goods:%d", uid)
	data, err := s.redis.Get(context.Background(), key).Result()
	if err == nil {
		if decodeErr := json.Unmarshal([]byte(data), &goods); decodeErr == nil {
			return goods, nil
		} else {
			zap.L().Warn("decode cached goods failed",
				zap.Uint("goods_id", uid),
				zap.String("trace_id", c.GetString(`trace_id`)),
				zap.Error(decodeErr),
			)
		}
	}
	goods, err = s.One(c, uid)
	if err != nil {
		fields := []zap.Field{
			zap.Uint("goods_id", uid),
			zap.String("trace_id", c.GetString(`trace_id`)),
		}
		zap.L().Warn("repo find goods failed", fields...)
		err = fmt.Errorf("service GetById:%w", err)
		return
	}
	btdata, _ := json.Marshal(goods)
	s.redis.Set(context.Background(), key, btdata, time.Hour)
	return
}
