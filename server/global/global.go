package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"xorm.io/xorm"

	"github.com/system-server2025/global/config"
)

var (
	LogFile *os.File
	Config  *config.Config
)



type Application struct {
	Echo   *echo.Echo
	Logger *logrus.Logger
	Redis  *redis.Client
	Xorm   *xorm.Engine
}



