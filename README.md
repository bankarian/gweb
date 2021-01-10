# 功能总结

## 1、http 服务

Engine 类型作为 gee 的核心，实现了 http.Handler 接口，用于提供基本的 http 服务。

## 2、路由

### 2.1 基础路由

以 map 来存储路由和与其绑定的处理函数，**URL-http 方法**作为 map 的 Key，具体的处理函数 handler 作为 map 的 Value。

### 2.2 路由组

路由组能更加方便地支持从 URL 的某一个部分分叉，拓展路由。

```go
// Using route group
engine := gee.New()
g := engine.Group("/somepath")
{
    g.GET("/aaa", /* ... */)
    g.GET("/bbb", /* ... */)
}
// Using common route
engine := gee.New()
g.GET("/sompath/aaa", /* ... */)
g.GET("/somepath/bbb", /* ... */)
```

为了支持路由组，所有的路由使用 Trie 来进行存储，**每一个 http 方法对应一棵 Trie**。

```go
type router struct {
	// method : *node
	roots map[string]*node
	// method-pattern : handler
	handlers map[string]HandlerFunc
}
```

```go
// Trie 节点类型
type node struct {
	pattern  string		// 完整的 URL 模式，可以携带参数。存储在 #叶子
	part     string		// 通过 "/" 分割出的每一个路径片段，存储在 #内部节点，用于索引到最终的 #叶子
	children []*node	// 当前节点的所有孩子
	isParam  bool		// 当前的 part 是否是参数。参数有两种
    							// 1. 变量参数 :param
    							// 2. 路径参数（只能包含一个） *path
}
```

## 3、上下文 Context

上下文的出现主要是统一处理 http 请求函数 handler 的结构，http 服务的所有参数都集合到结构体 Context 中，这样所有handler 结构就可以统一定义为一个 HandlerFunc。

```go
type Context struct {
	// origin objects
	Writer http.ResponseWriter  // 在 ServeHTTP 方法中将其初始化
	Req    *http.Request		// 在 ServeHTTP 方法中将其初始化

	// request info
	Path   string
	Method string
	Params map[string]string

	// response info
	StatusCode int

	// middlewares
	handlers []HandlerFunc
	index    int
}
```

```go
// HandlerFunc defines the request handler used by Gee
type HandlerFunc func(*Context)
```

## 4、中间件

中间件就是在处理服务端真正提供服务之前进行预处理，本质上也可以定义为一个 HandlerFunc，所以需要维护一个 handler 列表，将中间件存储在其中。那么当接收到一个 http 请求后，对应的 handler 添加到 handler 列表中，然后从头顺序执行整个 handler 列表，就达到了先执行中间件再真正处理请求的目的。

中间件列表的维护，由 Context 和 RouterGroup 完成。

<img src="https://gitee.com/bankarian/picStorage/raw/master/20210109175611.png" width="70%" />

