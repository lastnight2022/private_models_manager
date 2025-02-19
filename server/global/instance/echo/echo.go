package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	customMiddleware "github.com/system-server2025/middleware"
)

func InitEcho() *echo.Echo {
	E := echo.New()
	// 使用echo中间件
	E.Use(middleware.Logger(), middleware.Recover(), middleware.Gzip(),middleware.CORS())
	// 使用自定义中间件
	E.Use(customMiddleware.JWTMiddleware)
	// 静态文件
	E.Use(middleware.Static("static"))
	// 限制请求体大小
	E.Use(middleware.BodyLimit("2M"))

	return E
}