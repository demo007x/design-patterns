package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		fmt.Print(i)
		go getInstance()
		// go GetInstance()
	}
	fmt.Scanln()
}
