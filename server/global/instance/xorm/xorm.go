package xorm

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/system-server2025/global/config"
	"xorm.io/core"
	"xorm.io/xorm"
)


func ConnectDB(cfg config.Config) (*xorm.Engine,error) {
	var err error
	port := fmt.Sprintf("%d", cfg.Database.Port)
	var dsn = cfg.Database.User + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + port + ")/" + cfg.Database.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("dsn: ", dsn)
	Engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("初始化xorm数据库连接引擎失败: ", err)
		return nil,err
	}
	fmt.Println("succeed to connect to mysql")
	Engine.Logger().ShowSQL(true)
	Engine.SetMapper(core.GonicMapper{})
	return Engine,nil
}

func CloseDB(engine *xorm.Engine) {
	engine.Close()
}


