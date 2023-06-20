# Go **备忘录**模式讲解和代码示例

**备忘录**是一种行为设计模式， 允许生成对象状态的快照并在以后将其还原。

备忘录不会影响它所处理的对象的内部结构， 也不会影响快照中保存的数据。

## 概念示例

备忘录模式让我们可以保存对象状态的快照。 你可使用这些快照来将对象恢复到之前的状态。 这在需要在对象上实现撤销-重做操作时非常实用。

####  **originator.go:** 原发器

```go
package main

type Originator struct {
	state string
}

// 创建快照
func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

// 回复快照
func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) setState(state string) {
	e.state = state
}

func (e *Originator) getState() string {
	return e.state
}

```

####  **memento.go:** 备忘录

```go
package main

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}

```

####  **caretaker.go:** 负责人

```go
package main

type Caretaker struct {
	mementoArray []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	return c.mementoArray[index]
}

```

####  **main.go:** 客户端代码

```go
package main

import "fmt"

func main() {
	caretaker := &Caretaker{
		mementoArray: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	// 添加一个备忘录
	caretaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}

```

####  **output.txt:** 执行结果

```tex
Originator Current State: A
Originator Current State: C
Restored to State: C
Restored to State: A
```