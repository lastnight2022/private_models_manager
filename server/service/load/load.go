package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/system-server2025/global"
	"github.com/system-server2025/global/config"
	"github.com/system-server2025/global/instance/echo"
	"github.com/system-server2025/global/instance/logrus"
	"github.com/system-server2025/global/instance/redis"
	"github.com/system-server2025/global/instance/xorm"
)


func InitApp() *global.Application {
	var config config.Config
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("打开配置文件出错:", err)
		return nil
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("解析配置文件出错:", err)
		return nil
	}
	global.Config = &config
	fmt.Println(global.Config)
	var app *global.Application

	echo := echo.InitEcho()
	xorm,err := xorm.ConnectDB()
	if err != nil {
		fmt.Println("xorm 连接数据库失败")
	}
	redis := redis.ConnectRedis()
	logger := logrus.InitLogger()
	
	app = &global.Application{
		Echo: echo,
		Xorm: xorm,
		Redis: redis,
		Logger: logger,
	}
		
	fmt.Println("数据库连接字符串:", global.Config.Database.DBName)
	fmt.Println("服务器端口:", global.Config.Server.Port)
	return app
}
