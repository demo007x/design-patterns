# Go **工厂方法**模式讲解和代码示例

**工厂方法**是一种创建型设计模式， 解决了在不指定具体类的情况下创建产品对象的问题。

工厂方法定义了一个方法， 且必须使用该方法代替通过直接调用构造函数来创建对象 （ `new`操作符） 的方式。 子类可重写该方法来更改将被创建的对象所属类。

> 如果你不清楚**工厂**、 **工厂方法**和**抽象工厂**模式之间的区别， 请参阅[工厂模式比较](https://refactoringguru.cn/design-patterns/factory-comparison)。

## 概念示例

由于 Go 中缺少类和继承等 OOP 特性， 所以无法使用 Go 来实现经典的工厂方法模式。 不过， 我们仍然能实现模式的基础版本， 即简单工厂。

在本例中， 我们将使用工厂结构体来构建多种类型的武器。

首先， 我们来创建一个名为 `i­Gun`的接口， 其中将定义一支枪所需具备的所有方法。 然后是实现了 iGun 接口的 `gun`枪支结构体类型。 两种具体的枪支——`ak47`与 `musket`火枪 ——两者都嵌入了枪支结构体， 且间接实现了所有的 `i­Gun`方法。

`gun­Factory`枪支工厂结构体将发挥工厂的作用， 即通过传入参数构建所需类型的枪支。 *main.go* 则扮演着客户端的角色。 其不会直接与 `ak47`或 `musket`进行互动， 而是依靠 `gun­Factory`来创建多种枪支的实例， 仅使用字符参数来控制生产。

####  **iGun.go:** 产品接口

```go
package main

type iGun interface {
	setName(name string)
	setPower(power string)
	getName() string
	getPower() string
}

```

####  **gun.go:** 具体产品

```go
package main

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}


```

####  **ak47.go:** 具体产品

```go
package main

type Ak47 struct {
	Gun
}

func newAk47() iGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

```

####  **musket.go:** 具体产品

```go
package main

type musket struct {
	Gun
}

func newMusket() iGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

```

####  **gunFactory.go:** 工厂

```go
package main

import "fmt"

func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "nusket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}

```

####  **main.go:** 客户端代码

```

```

####  **output.txt:** 执行结果

```

```