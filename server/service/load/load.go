package service

import (
	"fmt"

	"github.com/system-server2025/global"
	"github.com/system-server2025/global/config"
	_ "github.com/system-server2025/global/instance/cron"
	"github.com/system-server2025/global/instance/echo"
	"github.com/system-server2025/global/instance/logrus"
	"github.com/system-server2025/global/instance/mongodb"
	"github.com/system-server2025/global/instance/redis"
	"github.com/system-server2025/global/instance/xorm"
)


func InitApp() *global.Application {
	config.Init("config.json")
	cfg := config.GetConfig()
	var app *global.Application

	echo := echo.InitEcho(cfg)
	xorm,err := xorm.ConnectDB(cfg)
	if err != nil {
		fmt.Println("xorm 连接数据库失败")
	}
	redis := redis.ConnectRedis(cfg)
	logger := logrus.InitLogger()
	mongo := mongodb.InitMongo(cfg)
	xorm.SetLogger(logger)
	
	app = &global.Application{
		Echo: echo,
		Xorm: xorm,
		Redis: redis,
		Logger: logger,
		Mongo: mongo,
	}
		
	fmt.Println("数据库连接字符串:", cfg.Database.DBName)
	fmt.Println("服务器端口:", cfg.Server.Port)
	return app
}
