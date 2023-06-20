package main

import "fmt"

func main() {
	hpPrinter := &HP{}
	esponPrinter := &Epson{}

	macComputer := &Mac{}
	winComputer := &Windows{}

	macComputer.SetPrinter(esponPrinter)
	macComputer.Print()
	fmt.Println()
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()
	winComputer.SetPrinter(esponPrinter)
	winComputer.Print()
	fmt.Println()
}
