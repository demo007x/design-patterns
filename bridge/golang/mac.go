package main

import "fmt"

type Mac struct {
	printer Printer
}

func (mac *Mac) Print() {
	fmt.Println("print request for mac ")
	mac.printer.PrintFile()
}

func (mac *Mac) SetPrinter(p Printer) {
	mac.printer = p
}
