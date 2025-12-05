package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	Client *redis.Client
	ctx    = context.Background()
)

// Init 初始化 Redis 连接
func Init() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	// 测试连接
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	return nil
}

// Get 获取缓存
func Get(key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

// Set 设置缓存
func Set(key string, value interface{}, expiration time.Duration) error {
	return Client.Set(ctx, key, value, expiration).Err()
}

// Del 删除缓存
func Del(keys ...string) error {
	return Client.Del(ctx, keys...).Err()
}

// Exists 检查键是否存在
func Exists(keys ...string) (int64, error) {
	return Client.Exists(ctx, keys...).Result()
}

// Incr 自增
func Incr(key string) (int64, error) {
	return Client.Incr(ctx, key).Result()
}

// Expire 设置过期时间
func Expire(key string, expiration time.Duration) error {
	return Client.Expire(ctx, key, expiration).Err()
}

// Close 关闭 Redis 连接
func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}

// GetClient 获取 Redis 客户端实例
func GetClient() *redis.Client {
	return Client
}
