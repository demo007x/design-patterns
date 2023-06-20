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
