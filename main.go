// file path: terminus/main.go
package main

import (
	"fmt"
	"terminus/internal/entity"
)


func main() {
	fmt.Println("Welcome to Terminus")
	director := entity.NewDirector(entity.NewEntityBuilder())
	entity := entity.NewEntityBuilder().
		SetName("Bob").
		SetLevel(1).
		Build()
	entityJson, _ := entity.ToJSON()
	fmt.Println(string(entityJson))
	warrior := director.BuildWarrior("Bob The Warrior")
	warriorJson, _ := warrior.ToJSON()
	fmt.Println(string(warriorJson))
}