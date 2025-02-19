package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	jwt "github.com/lastnight2022-cc/tools_lib/utils/jwt"
	"github.com/system-server2025/global"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // 获取请求路径
        requestPath := c.Request().URL.Path
        // 定义不需要验证的路径列表
        skipPaths := []string{"/login","/static"}

        // 检查是否需要跳过此路径
        for _, path := range skipPaths {
            if requestPath == path {
                return next(c)
            }
        }
        token := c.Request().Header.Get("Authorization")
        if token == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
        }
        parts := strings.Split(token, " ")
        if len(parts)!= 2 || parts[0]!= "Bearer" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
        }
        jwtToken := parts[1]
        _, err := jwt.VerifyJWTToken(jwtToken, global.GVA.Config.Server.Secret)
        if err!= nil {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
        }
        return next(c)
    }
}