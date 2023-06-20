package main

import "fmt"

type Customer struct {
	id string
}

func (c *Customer) getID() string {
	return c.id
}

func (c *Customer) update(iteName string) {
	fmt.Printf("Sendint email to customer %s for item %s\n", c.id, iteName)
}
