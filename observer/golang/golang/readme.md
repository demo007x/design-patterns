# Go **观察者**模式讲解和代码示例

**观察者**是一种行为设计模式， 允许一个对象将其状态的改变通知其他对象

观察者模式提供了一种作用于任何实现了订阅者接口的对象的机制， 可对其事件进行订阅和取消订阅。


## 概念示例

在电商网站中， 商品时不时地会出现缺货情况。 可能会有客户对于缺货的特定商品表现出兴趣。 这一问题有三种解决方案：

1. 客户以一定的频率查看商品的可用性。
2. 电商网站向客户发送有库存的所有新商品。
3. 客户只订阅其感兴趣的特定商品， 商品可用时便会收到通知。 同时， 多名客户也可订阅同一款产品。

选项 3 是最具可行性的， 这其实就是观察者模式的思想。 观察者模式的主要组成部分有：

- 会在有任何事发生时发布事件的主体。
- 订阅了主体事件并会在事件发生时收到通知的观察者。

####  **subject.go:** 主体

```go
package main

type Subject interface {
	register(observer Observer)
	deregister(obs Observer)
	notifyAll()
}
```

####  **item.go:** 具体主体

```go
package main

import (
	"fmt"
)

type Item struct {
	observerList []Observer // 多个观察者
	name         string
	inStock      bool // 有库存
}

func newItem(name string) *Item {
	return &Item{
		name:    name,
		inStock: false,
	}
}

// 更新状态
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock \n", i.name)
	i.inStock = true // 更新状态，有库存 s
	i.notifyAll()
}

func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *Item) notifyAll() {
	for _, v := range i.observerList {
		v.update(i.name)
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

```

####  **observer.go:** 观察者

```go
package main

type Observer interface {
	update(string)
	getID() string
}

```

####  **customer.go:** 具体观察者

```go
package main

import "fmt"

type Customer struct {
	id string
}

func (c *Customer) getID() string {
	return c.id
}

func (c *Customer) update(iteName string) {
	fmt.Printf("Sendint email to customer %s for item %s\n", c.id, iteName)
}

```

####  **main.go:** 客户端代码

```go
package main

func main() {
	shirtItem := newItem("Nick Shirt")
	observerFirst := &Customer{
		id: "abc@gmail.com",
	}

	observerSecond := &Customer{
		id: "def@gmail.com",
	}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}

```

####  **output.txt:** 执行结果

```text
Item Nick Shirt is now in stock 
Sendint email to customer abc@gmail.com for item Nick Shirt
Sendint email to customer def@gmail.com for item Nick Shirt
```