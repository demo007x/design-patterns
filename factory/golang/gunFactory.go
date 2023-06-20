package main

import "fmt"

func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "nusket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}
