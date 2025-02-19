package global

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"xorm.io/xorm"

	"github.com/system-server2025/global/config"
)

var (
	GVA     *GlobalValue
	LogFile *os.File
)

type GlobalValue struct {
	Config config.Config `json:"global"` // 全局配置
	App    *Application
}

type Application struct {
	echo   *echo.Echo
	es     *elasticsearch.Client
	logger *logrus.Logger
	redis  *redis.Client
	xorm   *xorm.Engine
}

func NewGlobalValue(echo *echo.Echo, es *elasticsearch.Client, logger *logrus.Logger, redis *redis.Client, xorm *xorm.Engine) *GlobalValue {
	return &GlobalValue{
		Config: config.Config{},
		App: &Application{
			echo:   echo,
			es:     es,
			logger: logger,
			redis:  redis,
			xorm:   xorm,
		},
	}
}
