package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/system-server2025/global"
	"github.com/system-server2025/global/instance/echo"
	elasticsearch "github.com/system-server2025/global/instance/elasticserach"
	"github.com/system-server2025/global/instance/logrus"
	"github.com/system-server2025/global/instance/redis"
	"github.com/system-server2025/global/instance/xorm"
)


func init() {
	echo := echo.InitEcho()
	xorm,err := xorm.ConnectDB()
	if err != nil {
		fmt.Println("xorm 连接数据库失败")
	}
	redis := redis.ConnectRedis()
	logger := logrus.InitLogger()
	es := elasticsearch.InitEs()
	global.GVA = global.NewGlobalValue(echo,es,logger,redis,xorm)
	
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("打开配置文件出错:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&global.GVA.Config)
	if err != nil {
		fmt.Println("解析配置文件出错:", err)
		return
	}
	fmt.Println("数据库连接字符串:", global.GVA.Config.Database.DBName)
	fmt.Println("服务器端口:", global.GVA.Config.Server.Port)
}
