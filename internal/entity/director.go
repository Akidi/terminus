// import path: terminus/internal/entity
// file path: ./internal/entity/director.go
package entity

import "terminus/internal/stats"

type ArchType string

const (
	None				ArchType = ""
	Adventurer	ArchType = "Adventurer"
	Warrior			ArchType = "Warrior"
	Mage				ArchType = "Mage"
	BattleMage	ArchType = "BattleMage"
	Rogue				ArchType = "Rogue"
	Cleric			ArchType = "Cleric"
)

type Director struct {
	builder EntityBuilder
}

func NewDirector(builder EntityBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) BuildBasic(name string) EntityBuilder {
	return d.builder.
		Reset().
		SetName(name).
		SetLevel(1)
}

func (d *Director) BuildWarrior(name string) Entity {
	return d.BuildBasic(name).
		SetArchType(Warrior).
		SetAttributes(stats.NewWarriorAttributes()).
		Build()
}

func (d *Director) BuildMage(name string) Entity {
	return d.BuildBasic(name).
		SetArchType(Mage).
		SetAttributes(stats.NewMageAttributes()).
		Build()
}

func (d *Director) BuildBattleMage(name string) Entity {
	return d.BuildBasic(name).
		SetArchType(Mage).
		SetAttributes(stats.NewBattleMageAttributes()).
		Build()
}

func (d *Director) BuildRogue(name string) Entity {
	return d.BuildBasic(name).
		SetArchType(Mage).
		SetAttributes(stats.NewRogueAttributes()).
		Build()
}

func (d *Director) BuildCleric(name string) Entity {
	return d.BuildBasic(name).
		SetArchType(Mage).
		SetAttributes(stats.NewClericAttributes()).
		Build()
}