package main

import "fmt"

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}
