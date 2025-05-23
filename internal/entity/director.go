// import path: terminus/internal/entity
// file path: ./internal/entity/director.go
package entity

import (
	"terminus/internal/modifiers"
	"terminus/internal/stats"
)

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

func (d *Director) buildBasic(name string) EntityBuilder {
	return d.builder.
		Reset().
		SetName(name).
		SetLevel(1)
}

func (d *Director) BuildWarrior(name string) Entity {
	return d.buildBasic(name).
		SetArchType(Warrior).
		SetAttributes(stats.NewWarriorAttributes()).
		Build()
}

func (d *Director) BuildMage(name string) Entity {
	archType := d.buildBasic(name).
		SetArchType(Mage).
		SetAttributes(stats.NewMageAttributes()).
		Build()

	archType.Attributes().
		Get("Strength").
		Modifiers().
		AddModifier(modifiers.NewModifier().
			SetName("Strength + 15").
			SetKind(modifiers.Flat).
			SetSource(modifiers.SourcePassive).
			SetTarget(modifiers.TargetStrength).
			SetValue(15).Build()).
		AddModifier(modifiers.NewModifier().
			SetName("Strength * 1.5").
			SetKind(modifiers.Multiplier).
			SetSource(modifiers.SourcePassive).
			SetTarget(modifiers.TargetStrength).
			SetValue(1.5).Build()).
		AddModifier(modifiers.NewModifier().
			SetName("Strength + 25%").
			SetKind(modifiers.Percent).
			SetSource(modifiers.SourcePassive).
			SetTarget(modifiers.TargetStrength).
			SetValue(0.25).Build())
		archType.Attributes().
		Get("Intelligence").
		Modifiers().
		AddModifier(modifiers.NewModifier().
			SetName("Intelligence + 1000").
			SetKind(modifiers.Flat).
			SetSource(modifiers.SourcePassive).
			SetTarget(modifiers.TargetIntelligence).
			SetValue(1250).Build()).
		AddModifier(modifiers.NewModifier().
			SetName("Intelligence * 1.75").
			SetKind(modifiers.Multiplier).
			SetSource(modifiers.SourcePassive).
			SetTarget(modifiers.TargetIntelligence).
			SetValue(1.75).Build())
		
	archType.Attributes().ApplyAll()
	return archType
}

func (d *Director) BuildBattleMage(name string) Entity {
	return d.buildBasic(name).
		SetArchType(Mage).
		SetAttributes(stats.NewBattleMageAttributes()).
		Build()
}

func (d *Director) BuildRogue(name string) Entity {
	return d.buildBasic(name).
		SetArchType(Mage).
		SetAttributes(stats.NewRogueAttributes()).
		Build()
}

func (d *Director) BuildCleric(name string) Entity {
	return d.buildBasic(name).
		SetArchType(Mage).
		SetAttributes(stats.NewClericAttributes()).
		Build()
}