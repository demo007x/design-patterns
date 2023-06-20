# Go **装饰**模式讲解和代码示例

**装饰**是一种结构设计模式， 允许你通过将对象放入特殊封装对象中来为原对象增加新的行为。

由于目标对象和装饰器遵循同一接口， 因此你可用装饰来对对象进行无限次的封装。 结果对象将获得所有封装器叠加而来的行为。

## 概念示例

####  **pizza.go:** 零件接口

```go
package main

type IPizza interface {
	getPrice() int
}

```

####  **veggieMania.go:** 具体零件

```go
package main

type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15
}

```

####  **tomatoTopping.go:** 具体装饰

```go
package main

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
```

####  **cheeseTopping.go:** 具体装饰

```go
package main

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

```

####  **main.go:** 客户端代码

```go
package main

import "fmt"

func main() {
	pizza := &VeggieMania{}
	pizzaWithCheese := &CheeseTopping{pizza: pizza}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}

```

####  **output.txt:** 执行结果

```

```