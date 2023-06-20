package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(buildType string) IBuilder {
	switch buildType {
	case "normal":
		return newNormalBuilder()
	case "igloo":
		return newIglooBuilder()
	}
	return nil
}
