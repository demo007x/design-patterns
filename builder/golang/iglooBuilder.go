package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "snow window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "snow door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		windowsType: b.windowType,
		doorType:    b.doorType,
		floor:       b.floor,
	}
}
