# Go **享元**模式讲解和代码示例

**享元**是一种结构型设计模式， 它允许你在消耗少量内存的情况下支持大量对象。

模式通过共享多个对象的部分状态来实现上述功能。 换句话来说， 享元会将不同对象的相同数据进行缓存以节省内存。

## 概念示例

在游戏 《反恐精英》 中， 恐怖分子和反恐精英身着不同类型的衣物。 为了简便起见， 我们就假设双方都各有一种服装类型。 服装对象嵌入在玩家对象之中， 如下所示。

下面是玩家的结构体。 我们可以看到， 服装对象是嵌入在玩家结构体之中的：

```go
type player struct {
  dress      dress
  playerType string // 可为 T 或 CT
  lat        int
  long       int
}
```

假设目前有 5 名恐怖分子和 5 名反恐精英， 一共是 10 名玩家。 那么关于服装， 我们就有两个选项了。

1. 10 个玩家对象各自创建不同的服装对象， 并将其嵌入。 总共会创建 10 个服装对象。
2. 我们创建两个服装对象：

- 单一恐怖分子服装对象： 其将在 5 名恐怖分子之间共享。
- 单一反恐精英服装对象： 其将在 5 名反恐精英之间共享。

你可以看到， 方法 1 中我们总共创建了 10 个服装对象； 方法 2 中则只有 2 个服装对象。 第二种方法， 就是我们所遵循的享元设计模式。 我们所创建的 2 个服装对象被称为是享元对象。

享元模式会从对象中提取出公共部分并创建享元对象。 这些享元对象 （服装） 随后可在多个对象 （玩家） 中分享。 这极大地减少了服装对象的数量， 更棒的是即便你创建了更多玩家， 也只需这么两个服装对象就足够了。

在享元模式中， 我们会将享元对象存储在 map 容器中。 每当创建共享享元对象的其他对象时， 都会从 map 容器中获取享元对象。

下面让我们来看看此类安排的内部状态和外部状态：

- *内部状态*： 内部状态的服装可在多个恐怖分子和反恐精英对象间共享。
- *外部状态*： 玩家位置和玩家所使用的武器就是外部状态， 因为其在每个对象中都是不同的。

####  **dressFactory.go:** 享元工厂

```go
package main

import "fmt"

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressFactoryByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerrorisDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.dressMap[dressType] = newCounterTerrroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

```

####  **dress.go:** 享元接口

```go
package main

// 享元接口
type Dress interface {
	getColor() string
}

```

####  **terroristDress.go:** 具体享元对象

```go
package main

type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerrorisDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

```

####  **counterTerroristDress.go:** 具体享元对象

```go
package main

// 享元接口
type CounterTerrroristDress struct {
	color string
}

func (c *CounterTerrroristDress) getColor() string {
	return c.color
}

func newCounterTerrroristDress() *CounterTerrroristDress {
	return &CounterTerrroristDress{color: "green"}
}

```

####  **player.go:** 背景

```go
package main

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressFactoryByType(dressType)
	return &Player{
		dress:      dress,
		playerType: playerType,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

```

####  **game.go:** 客户端代码

```go
package main

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 01),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}

```

####  **main.go:** 客户端代码

```go
package main

import "fmt"

func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}

```

####  **output.txt:** 执行结果

```
DressColorType: tDress
DressColor: red
DressColorType: ctDress
DressColor: green
```