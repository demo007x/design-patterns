# Go **策略**模式讲解和代码示例

**策略**是一种行为设计模式， 它将一组行为转换为对象， 并使其在原始上下文对象内部能够相互替换。

原始对象被称为上下文， 它包含指向策略对象的引用并将执行行为的任务分派给策略对象。 为了改变上下文完成其工作的方式， 其他对象可以使用另一个对象来替换当前链接的策略对象。

## 概念示例

思考一下构建内存缓存的情形。 由于处在内存中， 故其大小会存在限制。 在达到其上限后， 一些条目就必须被移除以留出空间。 此类操作可通过多种算法进行实现。 一些流行的算法有：

- 最少最近使用 （LRU）： 移除最近使用最少的一条条目。
- 先进先出 （FIFO）： 移除最早创建的条目。
- 最少使用 （LFU）： 移除使用频率最低一条条目。

问题在于如何将我们的缓存类与这些算法解耦， 以便在运行时更改算法。 此外， 在添加新算法时， 缓存类不应改变。

这就是策略模式发挥作用的场景。 可创建一系列的算法， 每个算法都有自己的类。 这些类中的每一个都遵循相同的接口， 这使得系列算法之间可以互换。 假设通用接口名称为 `eviction­Algo`移除算法 。

现在， 我们的主要缓存类将嵌入至 `eviction­Algo`接口中。 缓存类会将全部类型的移除算法委派给 `eviction­Algo`接口， 而不是自行实现。 鉴于 `eviction­Algo`是一个接口， 我们可在运行时将算法更改为 LRU、 FIFO 或者 LFU， 而不需要对缓存类做出任何更改。

####  **evictionAlgo.go:** 策略接口

```go
package main

// 策略接口
type EvictionAlgo interface {
	evict(c *Cache)
}

```

####  **fifo.go:** 具体策略

```go
package main

import "fmt"

type Fifo struct{}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

```

####  **lru.go:** 具体策略

```go
package main

import "fmt"

type Lru struct{}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

```

####  **lfu.go:** 具体策略

```go
package main

import "fmt"

type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}

```

####  **cache.go:** 背景

```go
package main

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	macCapacity  int
}

// 初始化的时候将策略注入到 cache 中
func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		macCapacity:  2,
	}
}

// 动态修改策略
func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	// 如果缓存中的容量等于了最大容量，则需要执行策略来移除 s
	if c.capacity == c.macCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

```

####  **main.go:** 客户端代码

```go
package main

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")
	lru := &Lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")
}

```

####  **output.txt:** 执行结果

```tcl
Evicting by lfu strtegy
Evicting by lru strtegy
Evicting by fifo strtegy
```