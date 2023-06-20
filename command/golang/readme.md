# Go **命令**模式讲解和代码示例

**命令**是一种行为设计模式， 它可将请求或简单操作转换为一个对象。

此类转换让你能够延迟进行或远程执行请求， 还可将其放入队列中。

## 概念示例

下面我们通过电视机的例子来了解命令模式。 你可通过以下方式打开电视机：

- 按下遥控器上的 ON 开关；
- 按下电视机上的 ON 开关。

我们可以从实现 ON 命令对象并以电视机作为接收者入手。 当在此命令上调用 `execute`执行方法时， 方法会调用 `TV.on`打开电视函数。 

最后的工作是定义请求者： 这里实际上有两个请求者： 遥控器和电视机。 两者都将嵌入 ON 命令对象。

注意我们是如何将相同请求封装进多个请求者的。 我们也可以采用相同的方式来处理其他命令。 创建独立命令对象的优势在于可将 UI 逻辑与底层业务逻辑解耦。 这样就无需为每个请求者开发不同的处理者了。 命令对象中包含执行所需的全部信息， 所以也可用于延迟执行。

####  **button.go:** 请求者

```go
package main

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

```

####  **command.go:** 命令接口

```go
package main

type Command interface {
	execute()
}
```

####  **onCommand.go:** 具体接口

```go
package main

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.On()
}

```

####  **offCommand.go:** 具体接口

```go
package main

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.Off()
}

```

####  **device.go:** 接收者接口

```go
package main

type Device interface {
	On()
	Off()
}
```

####  **tv.go:** 具体接收者

```go
package main

import "fmt"

type TV struct {
	isRunning bool
}

func (tv *TV) On() {
	tv.isRunning = true
	fmt.Println("Turning tv on")
}

func (tv *TV) Off() {
	tv.isRunning = false
	fmt.Println("Turning tv off")
}

```

####  **main.go:** 客户端代码

```go
package main

func main() {
	tv := &TV{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{device: tv}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}

	offButton.press()
}

```

####  **output.txt:** 执行结果

```shell
Turning tv on
Turning tv off
```

