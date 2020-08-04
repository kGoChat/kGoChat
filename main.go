package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"kGoChat/controller"
	"kGoChat/websocket"
)

func main() {
	var port = flag.String("port", "8080", "监听端口")
	var ip = flag.String("ip", "127.0.0.1", "监听ip")
	flag.Parse()

	fmt.Println("Hello, kGoChat!")

	app := iris.New()

	//app.Logger().SetLevel("debug")
	//（可选）添加两个内置处理程序（handlers）
	//可以从任何与 http 相关的恐慌（http-relative panics） 中恢复
	//并将请求记录到终端。
	app.Use(recover.New())
	app.Use(logger.New())

	// load templates.
	app.RegisterView(iris.HTML("./views", ".html"))
	//app.StaticWeb("/", "./public")
	app.StaticEmbedded("/", "./public", Asset, AssetNames)

	//  "/"  服务于一个基于根路由的控制器。
	mvc.New(app).Handle(new(controller.ExampleController))

	mvc.Configure(app.Party("/websocket"), websocket.ConfigureMVC)

	_ = app.Run(iris.Addr((*ip)+":"+(*port)), iris.WithoutServerError(iris.ErrServerClosed))
}
