package response

import "github.com/labstack/echo/v4"

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 500
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c echo.Context) {
	c.JSON(code,Response{
		data,
		msg,
	})
}

func Ok(c echo.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c echo.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c echo.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c echo.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c echo.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c echo.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c echo.Context) {
	Result(ERROR, data, message, c)
}
