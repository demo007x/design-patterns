# Go **中介者**模式讲解和代码示例

**中介者**是一种行为设计模式， 让程序组件通过特殊的中介者对象进行间接沟通， 达到减少组件之间依赖关系的目的。

中介者能使得程序更易于修改和扩展， 而且能更方便地对独立的组件进行复用， 因为它们不再依赖于很多其他的类。



## 概念示例

中介者模式的一个绝佳例子就是火车站交通系统。 两列火车互相之间从来不会就站台的空闲状态进行通信。  `station­Manager`车站经理可充当中介者， 让平台仅可由一列入场火车使用， 而将其他火车放入队列中等待。 离场火车会向车站发送通知， 便于队列中的下一列火车进站。

####  **train.go:** 组件

```go
package main

type Train interface {
	arrive()
	depart()
	permitArrival()
}

```

####  **passengerTrain.go:** 具体组件

```go
package main

import "fmt"

type PassengerTrain struct {
	mediator Mediator
}

// 火车停靠
func (pt *PassengerTrain) arrive() {
	if !pt.mediator.canArrive(pt) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: arrived")
}

// 获取离开
func (pt *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: leaving")
	pt.mediator.notifyAboutDeparture()
}

func (pt *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	pt.arrive()
}

```

####  **freightTrain.go:** 具体组件

```go
package main

import "fmt"

type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: arrived")
}
func (g *FreightTrain) depart() {
	fmt.Println("FreightTrain: leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.arrive()
}

```

####  **mediator.go:** 中介者接口

```go
package main

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
```

####  **stationManager.go:** 具体中介者

```go
package main

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

```

####  **main.go:** 客户端代码

```go
package main

func main() {
	stationManager := newStationManager()
	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}

```

####  **output.txt:** 执行结果

```tex

```