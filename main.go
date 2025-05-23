// file path: terminus/main.go
package main

import (
	"fmt"
	"terminus/internal/entity"
)


func main() {
	fmt.Println("Welcome to Terminus")
	director := entity.NewDirector(entity.NewEntityBuilder())
	warrior := director.BuildMage("Bob The Mage")
	fmt.Print(warrior.Attributes().String())
}