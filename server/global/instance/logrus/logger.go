package logrus

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/system-server2025/global"
)


func InitLogger() *logrus.Logger {
	Logger := logrus.New()
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	// 创建初始日志文件
	createLogFile()
	Logger.SetOutput(global.LogFile)
	return Logger
}

func createLogFile() {
	var err error
	// 确保logs目录存在，如果不存在则创建
	err = os.MkdirAll("logs", 0755)
	if err != nil {
		panic(err)
	}
	global.LogFile, err = os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
}

func RotateLogFile() {
	// 关闭当前日志文件
	global.LogFile.Close()
	// 获取当前日期作为归档文件名的一部分
	dateStr := time.Now().Format("20060102")
	// 重命名当前日志文件
	oldLogFileName := "logs/app.log"
	newLogFileName := filepath.Join("logs", "app."+dateStr+".log")
	os.Rename(oldLogFileName, newLogFileName)
	// 创建新的日志文件
	createLogFile()
}

func ArchiveLogs() {
	// 这里可以添加将旧日志文件移动到归档目录的逻辑，例如：
	// 假设归档目录是logs/archive
	archiveDir := "logs/archive"
	// 检查归档目录是否存在，不存在则创建
	if _, err := os.Stat(archiveDir); os.IsNotExist(err) {
		os.Mkdir(archiveDir, 0777)
	}
	// 遍历日志文件，将符合条件的文件移动到归档目录（这里只是示例，实际可能需要更复杂的条件）
	files, _ := filepath.Glob("logs/app*.log")
	for _, file := range files {
		// 这里可以添加更复杂的判断，如根据日期等条件
		newFilePath := filepath.Join(archiveDir, filepath.Base(file))
		os.Rename(file, newFilePath)
	}
}
