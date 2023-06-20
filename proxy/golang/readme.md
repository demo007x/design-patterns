# Go **代理**模式讲解和代码示例

**代理**是一种结构型设计模式， 让你能提供真实服务对象的替代品给客户端使用。 代理接收客户端的请求并进行一些处理 （访问控制和缓存等）， 然后再将请求传递给服务对象。

代理对象拥有和服务对象相同的接口， 这使得当其被传递给客户端时可与真实对象互换。

## 概念示例

Nginx 这样的 Web 服务器可充当应用程序服务器的代理：

- 提供了对应用程序服务器的受控访问权限。
- 可限制速度。
- 可缓存请求。

####  **server.go:** 主体

```go
package main

type server interface {
	handleRequest(string, string) (int, string)
}
```

####  **nginx.go:** 代理

```go
package main

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}
func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

```

####  **application.go:** 真实主体

```go
package main

type Application struct{}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "ok"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "user created"
	}
	return 404, "Not OK"
}

```

####  **main.go:** 客户端代码

```go
package main

import "fmt"

func main() {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"
	httpCode, msg := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, msg)

	httpCode, msg = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, msg)

	httpCode, msg = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, msg)

	httpCode, msg = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, msg)

	httpCode, msg = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, msg)
}

```

####  **output.txt:** 执行结果

```

Url: /app/status
HttpCode: 200
Body: ok

Url: /app/status
HttpCode: 200
Body: ok

Url: /app/status
HttpCode: 403
Body: Not Allowed

Url: /app/status
HttpCode: 201
Body: user created

Url: /app/status
HttpCode: 404
Body: Not OK
```