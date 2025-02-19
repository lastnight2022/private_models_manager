package xorm

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/system-server2025/global"
	"xorm.io/core"
	"xorm.io/xorm"
)


func ConnectDB() (*xorm.Engine,error) {
	var err error
	port := fmt.Sprintf("%d", global.GVA.Config.Database.Port)
	var dsn = global.GVA.Config.Database.User + ":" + global.GVA.Config.Database.Password + "@tcp(" + global.GVA.Config.Database.Host + ":" + port + ")/" + global.GVA.Config.Database.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
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


