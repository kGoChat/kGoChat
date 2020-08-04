# Go web Iris 开发笔记

## 编译状态 build status

go test : ![go test CI](https://github.com/kGoChat/kGoChat/workflows/Go/badge.svg)


## demo 测试

```shell script
# 打包静态文件
go get -u github.com/shuLhan/go-bindata/...
go-bindata ./public/...
# 不编译执行
go run .
# 编译后执行
go build -race -ldflags "-extldflags '-static'" -o kGoChat
./kGoChat
```

## 编译可执行文件

```shell script
# 编译的到 kGoChat 可执行文件
go build -race -ldflags "-extldflags '-static'"
```

## IRIS embedding打包静态文件

### 简介

此包将任何文件转换为可管理的Go源代码。用于将二进制数据嵌入到go程序中。在转换为原始字节切片之前，文件数据可选地进行gzip压缩。

它在go-bindata子目录中附带了一个命令行工具。此工具提供一组命令行选项，用于自定义生成的输出。

### 提示

- 先要安装`github.com/shuLhan/go-bindata/…`
- 执行`go-bindata ./public/…`会出现一个`bindata.go`文件
- 再行`main.go`

```shell script
go get -u github.com/shuLhan/go-bindata/...
go-bindata ./public/...
go build -race -ldflags "-extldflags '-static'"
```


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



