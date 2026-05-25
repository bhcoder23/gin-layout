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

var User *UserService

type UserService struct {
	redis *redis.Client
	interfaces.Service[*models.User]
}

func NewUserService(r *repos.UserRepo, redisClient *redis.Client) *UserService {
	return &UserService{
		redis:   redisClient,
		Service: *interfaces.NewService(repos.User),
	}
}

func init() {
	RegisterServices(func() {
		User = NewUserService(repos.User, components.Redis)
	})
}

func (s *UserService) GetByID(c *gin.Context) (user *models.User, err error) {
	uid := c.GetUint(`user_id`)
	key := fmt.Sprintf("user:%d", uid)
	data, err := components.Redis.Get(context.Background(), key).Result()
	if err == nil {
		if decodeErr := json.Unmarshal([]byte(data), &user); decodeErr == nil {
			return user, nil
		} else {
			zap.L().Warn("decode cached user failed",
				zap.Uint("user_id", uid),
				zap.String("trace_id", c.GetString(`trace_id`)),
				zap.Error(decodeErr),
			)
		}
	}
	user, err = repos.User.One(c, uid)
	if err != nil {
		fields := []zap.Field{
			zap.Uint("user_id", uid),
			zap.String("trace_id", c.GetString(`trace_id`)),
		}
		zap.L().Warn("repo find user failed", fields...)
		err = fmt.Errorf("service GetById:%w", err)
		return
	}
	btdata, _ := json.Marshal(user)
	components.Redis.Set(context.Background(), key, btdata, time.Hour)
	return
}
