package websocket

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"kGoChat/controller"
	"testing"
)

func TestWebsocketController_Get(t *testing.T) {
	app := iris.New()
	//  "/"  服务于一个基于根路由的控制器。
	mvc.New(app).Handle(new(controller.ExampleController))

	mvc.Configure(app.Party("/websocket"), ConfigureMVC)
	fmt.Print(iris.Addr("127.0.0.1:8132"), iris.WithoutServerError(iris.ErrServerClosed), "\n")
}
