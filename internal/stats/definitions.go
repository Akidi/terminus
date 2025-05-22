// import path: terminus/internal/stats
// file path: ./internal/stats/definitions.go
package stats

type AttrCategory int

const (
	Primary		AttrCategory = iota
	Derived
)

type AttributeDefinition interface {
	ID() 					string
	Name() 				string
	ShortName() 	string
	Description() string
	Category() 		AttrCategory
}

type AttributeDefinitionImpl struct {
	id					string
	name				string
	shortName		string
	description	string
	category		AttrCategory
}

func (ad *AttributeDefinitionImpl) ID() string {
	return ad.id
}

func (ad *AttributeDefinitionImpl) Name() string {
	return ad.name
}

func (ad *AttributeDefinitionImpl) ShortName() string {
	return ad.shortName
}

func (ad *AttributeDefinitionImpl) Description() string {
	return ad.description
}

func (ad *AttributeDefinitionImpl) Category() AttrCategory {
	return ad.category
}

var AttributeDefs = map[string]*AttributeDefinitionImpl{
	"Strength":     {id: "Strength", name: "Strength", shortName: "STR", description: "Physical power and carry capacity", category: Primary},
	"Dexterity":    {id: "Dexterity", name: "Dexterity", shortName: "DEX", description: "Speed, accuracy, evasion", category: Primary},
	"Vitality":     {id: "Vitality", name: "Vitality", shortName: "VIT", description: "Health, defense, endurance", category: Primary},
	"Intelligence": {id: "Intelligence", name: "Intelligence", shortName: "INT", description: "Magic damage and skill power", category: Primary},
	"Wisdom":       {id: "Wisdom", name: "Wisdom", shortName: "WIS", description: "Mana, resistance, healing", category: Primary},
	"Luck":         {id: "Luck", name: "Luck", shortName: "LCK", description: "Criticals, drops, randomness", category: Primary},
}