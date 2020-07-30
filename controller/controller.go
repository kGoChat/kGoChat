package controller

// ExampleController 服务于 "/", "/ping" and "/hello" 路由。
type ExampleController struct{}

// GetPing serves
// Method:   GET
// Resource: /ping
func (c *ExampleController) GetPing() string {
	return "pong"
}

// GetHelloIris serves
// Method:   GET
// Resource: /hello/iris
func (c *ExampleController) GetHelloIris() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

/*
// BeforeActivation 会被调用一次，在控制器适应主应用程序之前
// 并且当然也是在服务运行之前
// 在版本 9 之后，你还可以为特定控制器的方法添加自定义路由。
// 在这个控制器，你可以注册自定义方法的处理程序。
// 使用带有 `ca.Router` 的标准路由
// 在不适用 mvc 的情况下，你可以做任何你可以做的事情
// 并将添加的依赖项绑定到
// 一个控制器的字段或者方法函数的输入参数中。
func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}

	b.Handle(
		"GET",
		"/custom_path",
		"CustomHandlerWithoutFollowingTheNamingGuide",
		anyMiddlewareHere,
	)

	// 或者甚至可以添加基于这个控制器路由的全局中间件，
	// 比如在这个例子里面的根路由 "/" :
	// b.Router().Use(myMiddleware)
}

// CustomHandlerWithoutFollowingTheNamingGuide serves
// Method:   GET
// Resource: /custom_path
func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}
*/
// GetUserBy serves
// Method:   GET
// Resource: http://localhost:8080/user/{username:string}
// By 是一个保留的“关键字”来告诉框架你要
// 在函数的输入参数中绑定路径参数，它也是
// 有助于在同一个控制器中拥有 "Get“ 和 "GetBy"。
//
// func (c *ExampleController) GetUserBy(username string) mvc.Result {
//     return mvc.View{
//         Name: "user/username.html",
//         Data: username,
//     }
// }

/*
在一个控制器中可以使用多个 HTTP 方法，工厂会
为每条路由注册正确的 HTTP 方法，你可以按需取消注释：

func (c *ExampleController) Post() {}
func (c *ExampleController) Put() {}
func (c *ExampleController) Delete() {}
func (c *ExampleController) Connect() {}
func (c *ExampleController) Head() {}
func (c *ExampleController) Patch() {}
func (c *ExampleController) Options() {}
func (c *ExampleController) Trace() {}
*/

/*
func (c *ExampleController) All() {}
//        或者
func (c *ExampleController) Any() {}



func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
    // 1 -> HTTP 方法
    // 2 -> 路由路径
    // 3 -> 控制器的方法名称应该是该路由的处理程序（handlers）。
    b.Handle("GET", "/mypath/{param}", "DoIt", optionalMiddlewareHere...)
}

//激活后，所有依赖项都被设置 - 因此只能访问它们
//但仍可以添加自定义控制器或者简单的标准处理程序（handlers）。
func (c *ExampleController) AfterActivation(a mvc.AfterActivation) {}

*/
