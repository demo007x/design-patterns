# Go **单例**模式讲解和代码示例

**单例**是一种创建型设计模式， 让你能够保证一个类只有一个实例， 并提供一个访问该实例的全局节点。

单例拥有与全局变量相同的优缺点。 尽管它们非常有用， 但却会破坏代码的模块化特性。

在某些其他上下文中， 你不能使用依赖于单例的类。 你也将必须使用单例类。 绝大多数情况下， 该限制会在创建单元测试时出现。

## 概念示例

通常而言， 单例实例会在结构体首次初始化时创建。 为了实现这一操作， 我们在结构体中定义一个 `get­Instance`获取实例方法。 该方法将负责创建和返回单例实例。 创建后， 每次调用 `get­Instance`时都会返回相同的单例实例。

协程方面又有什么需要注意的吗？ 每当多个协程想要访问实例时， 单例结构体就必须返回相同的实例。 正因如此， 单例设计模式的实施工作很容易出错。 下方的例子表示了创建单例的正确方式。

一些值得注意的地方：

- 最开始时会有 `nil`检查， 确保 `single­Instance`单例实例在最开始时为空。 这是为了防止在每次调用 `get­Instance`方法时都去执行消耗巨大的锁定操作。 如果检查不通过， 则就意味着 `single­Instance`字段已被填充。
- `single­Instance`结构体将在锁定期间创建。
- 在获取到锁后还会有另一个 `nil`检查。 这是为了确保即便是有多个协程绕过了第一次检查， 也只能有一个可以创建单例实例。 否则， 所有协程都会创建自己的单例结构体实例。



####  **single.go:** 单例

```go
package main

import (
    "fmt"
    "sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating single instance now.")
            singleInstance = &single{}
        } else {
            fmt.Println("Single instance already created.")
        }
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}
```

####  **main.go:** 客户端代码

```go
package main

import (
    "fmt"
)

func main() {

    for i := 0; i < 30; i++ {
        go getInstance()
    }

    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}
```

####  **output.txt:** 执行结果

```apl
Creating single instance now.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
Single instance already created.
```



## 另一个例子

1.  `init`函数

我们可以在 `init`函数中创建单例实例。 这仅适用于实例的早期初始化工作已经确定时。  `init`函数仅会在包中的每个文件里调用一次， 所以我们可以确定其只会创建一个实例。

1.  `sync.Once`

`sync.Once`仅会执行一次操作。 可查看下面的代码：

#### **syncOnce.go:** 单例

```go
package main

import (
    "fmt"
    "sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        once.Do(
            func() {
                fmt.Println("Creating single instance now.")
                singleInstance = &single{}
            })
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}
```

#### **main.go:** 客户端代码

```go
package main

import (
    "fmt"
)

func main() {

    for i := 0; i < 30; i++ {
        go getInstance()
    }

    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}
```

#### **output.txt:** 执行结果

```api
Creating single instance now.
Single instance already created.
Single instance already created.
```

