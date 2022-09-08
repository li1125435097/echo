package main

import (
	"github.com/labstack/echo/v4"
	user "echo/src/controller/user"
	index "echo/src/controller/index"
	"reflect"
	"fmt"
)

const PORT = ":3000"
var typeof = reflect.TypeOf
var l = fmt.Println 

func main() {
	// 创建服务器
	server := echo.New()

	// 路由注册
	index.Init(server)
	user.Init(server)
	
	// 启动服务器
	server.Logger.Fatal(server.Start(PORT))
}
