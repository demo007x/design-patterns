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
