package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {
	// 定义一个简单的GET请求路由
	e.GET("/", func(c echo.Context) (err error) {
		return c.String(http.StatusOK, "gogogo")
	})
}
