package mongodb

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/system-server2025/global/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(cfg config.Config) *mongo.Client {
	// 连接配置
	// uri := "mongodb://prod_user:ProdPass123@10.0.0.5:27017,10.0.0.6:27017/admin?" +
	// 	"replicaSet=myReplicaSet&readPreference=secondaryPreferred"
	uri := buildMongoURI(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.MongoDb.ConnectTimeoutSec)*time.Second)
	defer cancel()

	// 创建带连接池的客户端
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(cfg.MongoDb.MaxPoolSize).
		SetMinPoolSize(cfg.MongoDb.MinPoolSize).
		SetSocketTimeout(time.Duration(cfg.MongoDb.SocketTimeoutSec)*time.Second),
	)
	if err != nil {
		log.Fatal("连接失败:", err)
	}

	// 健康检查
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Ping 失败:", err)
	}
	log.Println("成功连接到 MongoDB ")

	return client

}

func CloseMongo(client *mongo.Client) {
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := client.Disconnect(ctxShutdown); err != nil {
		log.Fatal("关闭连接失败:", err)
	}
	log.Println("连接已关闭")
}

func buildMongoURI(cfg config.Config) string {
	// 处理主机列表
	var hosts []string
	for _, h := range cfg.MongoDb.Hosts {
		hosts = append(hosts, fmt.Sprintf("%s:%d", h.IP, h.Port))
	}
	hostStr := strings.Join(hosts, ",")

	// 转义用户名和密码中的特殊字符
	escapedUser := url.PathEscape(cfg.MongoDb.Username)
	escapedPass := url.PathEscape(cfg.MongoDb.Password)

	// 构建基础 URI
	uri := ""
	if cfg.MongoDb.Username != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s/%s",
			escapedUser,
			escapedPass,
			hostStr,
			cfg.MongoDb.Database,
		)
	} else {
		uri = fmt.Sprintf("mongodb://%s/%s",
			hostStr,
			cfg.MongoDb.Database,
		)
	}

	// 添加查询参数
	queryParams := url.Values{}
	if cfg.MongoDb.AuthSource != "" {
		queryParams.Add("authSource", cfg.MongoDb.AuthSource)
	}
	if cfg.MongoDb.ReplicaSet != "" {
		queryParams.Add("replicaSet", cfg.MongoDb.ReplicaSet)
	}
	if cfg.MongoDb.ReadPreference != "" {
		queryParams.Add("readPreference", cfg.MongoDb.ReadPreference)
	}

	// 拼接完整 URI
	if len(queryParams) > 0 {
		uri += "?" + queryParams.Encode()
	}

	return uri
}
