// import path: terminus/internal/stats
// file path: ./internal/stats/stats.go
package stats

import (
	"fmt"
	"terminus/internal/common"
	"terminus/internal/modifiers"

	"github.com/charmbracelet/lipgloss"
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
	Modifiers()		modifiers.ModifiersImpl
	SetModified(modified float64)
	String()	string
	AsJSON() AttributeJson
}

type AttributeImpl struct {
	definition	AttributeDefinition
	current			float64
	modified		float64
	modifiers		modifiers.ModifiersImpl
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

func (a *AttributeImpl) Modifiers() modifiers.ModifiersImpl {
	return a.modifiers
}

var (
	attrStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("69"))
	baseStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	modifiedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
)

func (a *AttributeImpl) String() string {
	label := attrStyle.Render(fmt.Sprintf("%-12s", a.definition.Name()))
	base := baseStyle.Render(fmt.Sprintf("Base: %-4.0f", a.Current()))
	mod := modifiedStyle.Render(fmt.Sprintf("Modified: %-4.0f", a.Modified()))
	return fmt.Sprintf("%s  %s  %s", label, base, mod)
}

type AttributeJson struct {
	Name 			string	`json:"name"`
	Current		float64	`json:"current"`
	Modified	float64	`json:"modified"`
	Modifiers []modifiers.ModifierJson	`json:"modifiers"`
}

func (a *AttributeImpl) AsJSON() AttributeJson {
	return AttributeJson{
		Name: a.definition.Name(),
		Current: a.Current(),
		Modified: a.Modified(),
		Modifiers: a.modifiers.AsJSON(),
	}
}

func (a *AttributeImpl) ToJSON() ([]byte, error) {
	return common.ToJSON(a)
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
	String()					string
	ApplyAll()
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

var (
	titleStyle     = lipgloss.NewStyle().Bold(true).Underline(true).Foreground(lipgloss.Color("63"))
	sectionStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("219"))
	modListStyle   = lipgloss.NewStyle().MarginLeft(2)
	modEmptyStyle  = lipgloss.NewStyle().Faint(true).Italic(true).MarginLeft(2)
)

func (a AttributesImpl) String() string {
	var s string
	s += titleStyle.Render("Attributes:") + "\n\n"
	for _, attrName := range AttributeNames {
		attr := a.Get(attrName)
		s += attr.String() + "\n"
		s += attr.Modifiers().String() + "\n"
	}
	return s
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

func (a *AttributesImpl) ApplyAll() {
	for _, k := range AttributeNames {
		attr:=a.Get(k)
		attr.SetModified(attr.Modifiers().ApplyAll(attr.Current()))
	}
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