package global

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

var (
	LogFile *os.File
)

type Application struct {
	Echo   *echo.Echo
	Logger *logrus.Logger
	Redis  *redis.Client
	Xorm   *xorm.Engine
}
