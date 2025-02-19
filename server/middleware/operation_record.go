package middleware

import "github.com/labstack/echo/v4"

func OperationRecord(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

// package middleware

// import (
// 	"bytes"
// 	"github.com/kataras/iris/v12"
// 	"github.com/kataras/iris/v12/context"
// 	"go.uber.org/zap"
// 	"io"
// 	"net/http"
// 	"system-server/global"
// 	"system-server/model"
// 	"system-server/model/request"
// 	"system-server/service"
// 	"strconv"
// 	"time"
// )

// func OperationRecord() iris.Handler {
// 	return func(c iris.Context) {
// 		var body []byte
// 		var userId int
// 		if c.Request().Method != http.MethodGet {
// 			var err error
// 			body, err = io.ReadAll(c.Request().Body)
// 			if err != nil {
// 				global.IRA_LOG.Error("read body from request error:", zap.Any("err", err))
// 			} else {
// 				c.Request().Body = io.NopCloser(bytes.NewBuffer(body))
// 			}
// 		}
// 		if claims := c.Values().Get("claims"); claims != nil {
// 			waitUse := claims.(*request.CustomClaims)
// 			userId = int(waitUse.ID)
// 		} else {
// 			id, err := strconv.Atoi(c.Request().Header.Get("x-user-id"))
// 			if err != nil {
// 				userId = 0
// 			}
// 			userId = id
// 		}
// 		record := model.SysOperationRecord{
// 			Ip:     c.RemoteAddr(),
// 			Method: c.Request().Method,
// 			Path:   c.Request().URL.Path,
// 			Agent:  c.Request().UserAgent(),
// 			Body:   string(body),
// 			UserID: userId,
// 		}
// 		writer := &responseBodyWriter{
// 			ResponseWriter: c.ResponseWriter(),
// 			body:           &bytes.Buffer{},
// 		}
// 		c.ResetResponseWriter(writer)
// 		now := time.Now()

// 		c.Next()

// 		latency := time.Now().Sub(now)
// 		record.ErrorMessage = c.Values().GetString("errorMessage")
// 		record.Status = c.ResponseWriter().StatusCode()
// 		record.Latency = latency
// 		record.Resp = writer.body.String()

// 		if err := service.CreateSysOperationRecord(record); err != nil {
// 			global.IRA_LOG.Error("create operation record error:", zap.Any("err", err))
// 		}
// 	}
// }

// type responseBodyWriter struct {
// 	context.ResponseWriter
// 	body *bytes.Buffer
// }

// func (r responseBodyWriter) Write(b []byte) (int, error) {
// 	r.body.Write(b)
// 	return r.ResponseWriter.Write(b)
// }
