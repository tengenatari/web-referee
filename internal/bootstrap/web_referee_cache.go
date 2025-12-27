package bootstrap

import (
	"fmt"

	"github.com/tengenatari/web-referee/config"
	"github.com/tengenatari/web-referee/internal/storage/redisstorage"
)

func InitRedis(cfg *config.Config) *redisstorage.RedisStorage {
	cacheAddress := fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port)
	return redisstorage.NewRedisStorage(cacheAddress)
}
