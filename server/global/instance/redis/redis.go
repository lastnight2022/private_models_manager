package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/system-server2025/global"
)


func ConnectRedis() *redis.Client {
	// 读取Redis配置信息
	// 创建Redis客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     global.GVA.Config.RedisConfig.Addr,
		Password: global.GVA.Config.RedisConfig.Password,
		DB:       global.GVA.Config.RedisConfig.DB,
	})
	// 测试连接
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("连接Redis失败:", err)
		return nil
	}
	fmt.Println("连接Redis成功:", pong)
	return client
}


