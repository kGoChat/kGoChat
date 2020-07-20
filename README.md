# Go web Iris 开发笔记

## 注意事项

### 包问题

如果官方的 `v12` 有问题，则直接使用：
```go
import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
)
```

### Iris 路径问题

`Iris` 使用驼峰命名，每个大写字母换为 `/` + 小写字母，例如 `GetHelloIris` 的路径为 `GET Method` 的 `/hello/iris`

```go
// ExampleController 服务于路由。
type ExampleController struct{}

// Get serves
// Method:   GET
// Resource: /
func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

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
```



