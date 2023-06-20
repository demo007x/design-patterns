package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Create single instance new.")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("single instalce already create")
	}

	return singleInstance
}
