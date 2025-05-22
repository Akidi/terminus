// import path: terminus/internal/stats
// file path: ./internal/stats/stats.go
package stats

import (
	"terminus/internal/common"
	"terminus/internal/modifiers"
)

type AttributeName string

const (
	Strength     AttributeName = "Strength"
	Dexterity    AttributeName = "Dexterity"
	Vitality     AttributeName = "Vitality"
	Intelligence AttributeName = "Intelligence"
	Wisdom       AttributeName = "Wisdom"
	Luck         AttributeName = "Luck"
)

var AttributeNames = []AttributeName{
	Strength,
	Dexterity,
	Vitality,
	Intelligence,
	Wisdom,
	Luck,
}

type Attribute interface {
	Current()			float64
	SetCurrent(current float64)
	Modified()		float64
	SetModified(modified float64)
	AsJSON() AttributeJson
}

type AttributeImpl struct {
	definition	AttributeDefinition
	current			float64
	modified		float64
	modifiers		[]modifiers.Modifier
}

func (a *AttributeImpl) Current() float64 {
	return a.current
}

func (a *AttributeImpl) SetCurrent(current float64) {
	a.current = current
}

func (a *AttributeImpl) Modified() float64 {
	return a.modified
}

func (a *AttributeImpl) SetModified(modified float64) {
	a.modified = modified
}

func (a *AttributeImpl) Modifiers() []modifiers.Modifier {
	return a.modifiers
}

type AttributeJson struct {
	Name string	`json:"name"`
	Current	float64	`json:"current"`
	Modifiers []modifiers.ModifierJson	`json:"modifiers"`
}

func (a *AttributeImpl) AsJSON() AttributeJson {
	modifiers := make([]modifiers.ModifierJson, 0, len(a.Modifiers()))
	for _, mod := range a.modifiers {
		modifiers = append(modifiers, mod.AsJSON())
	}
	return AttributeJson{
		Name: a.definition.Name(),
		Current: a.Current(),
		Modifiers: modifiers,
	}
}

func (a *AttributeImpl) ToJSON() ([]byte, error) {
	return common.ToJSON(a)
}

func NewAttribute(name string, shortName string, description string, category AttrCategory, initialValue float64) Attribute {
	return &AttributeImpl{
		definition: &AttributeDefinitionImpl{
			name: name,
			shortName: shortName,
			description: description,
			category: category,
		},
		current: initialValue,
		modified: initialValue,
		modifiers: []modifiers.Modifier{},
	}
}

func NewAttributeFromDef(def *AttributeDefinitionImpl, value float64) Attribute {
	return &AttributeImpl{
		definition: def,
		current:    value,
		modified:   value,
		modifiers:  []modifiers.Modifier{},
	}
}

type Attributes interface {
	Strength()				Attribute
	Dexterity()				Attribute
	Vitality()				Attribute
	Intelligence()		Attribute
	Wisdom()					Attribute
	Get(name AttributeName)	Attribute
	All()							[]Attribute
	AsJSON()					[]AttributeJson
}

type AttributesImpl struct {
	attributes map[AttributeName]Attribute
}

func (a AttributesImpl) Strength() Attribute {
	return a.attributes["Strength"]
}

func (a AttributesImpl) Dexterity() Attribute {
	return a.attributes["Dexterity"]
}

func (a AttributesImpl) Vitality() Attribute {
	return a.attributes["Vitality"]
}

func (a AttributesImpl) Intelligence() Attribute {
	return a.attributes["Intelligence"]
}

func (a AttributesImpl) Wisdom() Attribute {
	return a.attributes["Wisdom"]
}

func (a AttributesImpl) Luck() Attribute {
	return a.attributes["Luck"]
}

func (a AttributesImpl) Get(name AttributeName) Attribute {
	return a.attributes[name]
}

func (a AttributesImpl) All() []Attribute {
	out := make([]Attribute, 0, len(a.attributes))
	for _, attr := range a.attributes {
		out = append(out, attr)
	}

	return out
}

type AttributesJson struct {
	Attributes []AttributeJson `json:"attributes"`
}

func (a *AttributesImpl) ToJSON() ([]byte, error) {
	return common.ToJSON(a)
}

func (a *AttributesImpl) AsJSON() []AttributeJson {
	attrs := make([]AttributeJson, 0, len(AttributeNames))
	for _, k := range AttributeNames {
		if attr, ok := a.attributes[k]; ok {
			attrs = append(attrs, attr.AsJSON())
		}
	}

	return attrs
}

func NewAttributes() Attributes {
	return &AttributesImpl{
		attributes: map[AttributeName]Attribute{
			"Strength": NewAttributeFromDef(AttributeDefs["Strength"], 1),
			"Dexterity": NewAttributeFromDef(AttributeDefs["Dexterity"], 1),
			"Vitality": NewAttributeFromDef(AttributeDefs["Vitality"], 1),
			"Intelligence": NewAttributeFromDef(AttributeDefs["Intelligence"], 1),
			"Wisdom": NewAttributeFromDef(AttributeDefs["Wisdom"], 1),
			"Luck": NewAttributeFromDef(AttributeDefs["Luck"], 1),
		},
	}
}

