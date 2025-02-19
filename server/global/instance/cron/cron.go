package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/system-server2025/global/instance/logrus"
)

func init() {
	// 初始化日志记录器（在util包的init函数中完成）
	// 这里假设已经在util包中初始化了Logger等相关变量

	// 创建一个cron实例
	c := cron.New()
	// 添加定时任务，每天凌晨0点切割日志文件
	_, err := c.AddFunc("0 0 0 * * *", func() {
		logrus.RotateLogFile()
		logrus.ArchiveLogs()
	})
	if err != nil {
		// logrus.Logger.Fatal("添加定时任务失败: ", err)
	}
	// 启动cron任务调度器
	c.Start()

	// 这里可以继续添加项目的其他初始化代码和主逻辑代码
	//...
}
