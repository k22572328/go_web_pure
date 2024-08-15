package main

import (
	"go_web_pure/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// 初始化路由
	routes.InitRoutes(e)

	// 启动服务器
	e.Logger.Fatal(e.Start(":8080"))
}
