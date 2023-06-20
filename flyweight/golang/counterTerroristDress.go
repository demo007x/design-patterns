package main

// 享元接口
type CounterTerrroristDress struct {
	color string
}

func (c *CounterTerrroristDress) getColor() string {
	return c.color
}

func newCounterTerrroristDress() *CounterTerrroristDress {
	return &CounterTerrroristDress{color: "green"}
}
