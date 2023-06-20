package main

import (
	"fmt"
	"log"
)

func main() {
	vendingMachine := newVendingMachine(1, 10)

	// 获取一个商品
	if err := vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vendingMachine.insertMoney(10); err != nil {
		log.Fatal(err.Error())
	}

	if err := vendingMachine.dispenseItem(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("================")
	if err := vendingMachine.addItem(2); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println()
	if err := vendingMachine.requestItem(); err != nil {
		log.Fatal(err.Error())
	}

	if err := vendingMachine.insertMoney(10); err != nil {
		log.Fatal(err.Error())
	}

	if err := vendingMachine.dispenseItem(); err != nil {
		log.Fatal(err.Error())
	}
}
