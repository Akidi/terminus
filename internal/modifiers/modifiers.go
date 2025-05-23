// import path: terminus/internal/modifier
// file path: ./internal/modifiers/modifier.go
package modifiers

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"terminus/internal/common"

	"github.com/charmbracelet/lipgloss"
)

type ModifierTarget string

const (
	TargetStrength     ModifierTarget = "Strength"
	TargetDexterity    ModifierTarget = "Dexterity"
	TargetVitality     ModifierTarget = "Vitality"
	TargetIntelligence ModifierTarget = "Intelligence"
	TargetWisdom       ModifierTarget = "Wisdom"
	TargetLuck         ModifierTarget = "Luck"

	TargetMaxHP ModifierTarget = "Max HP"
	TargetMaxMP ModifierTarget = "Max MP"
)

var ModifierTargetMap = map[string]ModifierTarget{
	"Strength": TargetStrength,
	"Dexterity": TargetDexterity,
	"Vitality": TargetVitality,
	"Intelligence": TargetIntelligence,
	"Wisdom": TargetWisdom,
	"Luck": TargetLuck,

	"HP_Max": TargetMaxHP,
	"MP_Max": TargetMaxMP,
}

type ModifierSourceType string

const (
	SourceItem    ModifierSourceType = "Item"
	SourceSkill   ModifierSourceType = "Skill"
	SourceBuff    ModifierSourceType = "Buff"
	SourceDebuff  ModifierSourceType = "Debuff"
	SourcePassive ModifierSourceType = "Passive"
)



type ModifierKind string

const (
	Flat       ModifierKind = "Flat"
	Multiplier ModifierKind = "Multiplier"
	Percent    ModifierKind = "Percent"
)

type Modifier interface {
	ID() 			string
	Name()		string
	Target()	ModifierTarget
	Kind()		ModifierKind
	Value()		float64
	Source()	ModifierSourceType
	Apply(base float64) float64
	AsJSON() ModifierJson
	String() string
}

type ModifierImpl struct {
	id     string
	name   string
	target ModifierTarget
	kind   ModifierKind
	value  float64
	source ModifierSourceType
}

func (m *ModifierImpl) ID() string {
	return m.id
}

func (m *ModifierImpl) Name() string {
	return m.name
}

func (m *ModifierImpl) Target() ModifierTarget {
	return m.target
}

func (m *ModifierImpl) Kind() ModifierKind {
	return m.kind
}

func (m *ModifierImpl) Value() float64 {
	return m.value
}

func (m *ModifierImpl) Source() ModifierSourceType {
	return m.source
}

func (m *ModifierImpl) Apply(base float64) float64 {
	switch m.kind {
	case Flat:
		return base + m.value
	case Percent:
		return base * (1 + m.value)
	case Multiplier:
		return base * m.value
	default:
		return base
	}
}

func (m *ModifierImpl) AsJSON() ModifierJson {
	return ModifierJson{
		ID:     m.id,
		Name:   m.name,
		Target: m.target,
		Kind:   m.kind,
		Value:  m.value,
		Source: m.source,
	}
}

func (m *ModifierImpl) ToJSON() ([]byte, error) {
	return common.ToJSON(m)
}

var (
	flatStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("213"))
	percentStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("141"))
	multiplierStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
	sourceStyle     = lipgloss.NewStyle().Faint(true)
	targetStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("69"))
)

func (m *ModifierImpl) String() string {
	target := targetStyle.Render(fmt.Sprintf("%-12s", m.Target()))

	switch m.Kind() {
	case Flat:
		val := flatStyle.Render(fmt.Sprintf("+%.0f", m.Value()))
		return fmt.Sprintf("%s  %s  %s", target, val, sourceStyle.Render(fmt.Sprintf("[from %s]", m.Source())))
	case Percent:
		val := percentStyle.Render(fmt.Sprintf("+%.0f%%", m.Value()*100))
		return fmt.Sprintf("%s  %s  %s", target, val, sourceStyle.Render(fmt.Sprintf("[from %s]", m.Source())))
	case Multiplier:
		val := multiplierStyle.Render(fmt.Sprintf("x%.2f", m.Value()))
		return fmt.Sprintf("%s  %s  %s", target, val, sourceStyle.Render(fmt.Sprintf("[from %s]", m.Source())))
	default:
		val := flatStyle.Render(fmt.Sprintf("%.0f", m.Value()))
		return fmt.Sprintf("%s  %s  %s", target, val, sourceStyle.Render(fmt.Sprintf("[from %s]", m.Source())))
	}
}

type ModifierJson struct {
	ID     string             `json:"id"`
	Name   string             `json:"name"`
	Target ModifierTarget     `json:"target"`
	Kind   ModifierKind       `json:"kind"`
	Value  float64            `json:"value"`
	Source ModifierSourceType `json:"source"`
}

type Modifiers interface {
	AddModifier(mod Modifier) Modifiers
	RemoveModifier(id string) Modifiers
	ApplyAll(base float64) float64
	ToJSON() ([]byte, error)
	AsJSON() []ModifierJson
	String()	string
}

type ModifiersImpl struct {
	modifiers map[ModifierTarget][]Modifier
}

func (m ModifiersImpl) AddModifier(mod Modifier) Modifiers {
	m.modifiers[mod.Target()] = append(m.modifiers[mod.Target()], mod)
	return m
}

func (m ModifiersImpl) RemoveModifier(id string) Modifiers {
	for source, mods := range m.modifiers {
		for i, mod := range mods {
			if mod.ID() == id {
				m.modifiers[source] = slices.Delete(mods, i, i+1)
			}
		}
	}
	return m
}

func (m ModifiersImpl) ApplyAll(base float64) float64 {
	for _, mods := range m.modifiers {
		sort.SliceStable(mods, func(i, j int) bool {
			return mods[i].Kind() < mods[j].Kind()
		})
		for _, mod := range mods {
			base = mod.Apply(base)
		}
	}

	return base
}

func (m ModifiersImpl) AsJSON() []ModifierJson {
	var out []ModifierJson
	for _, mods := range m.modifiers {
		for _, mod := range mods {
			out = append(out, mod.AsJSON())
		}
	}
	return out
}

func (m ModifiersImpl) ToJSON() ([]byte, error) {
	return common.ToJSON(m)
}

var (
	modEmptyStyle  = lipgloss.NewStyle().Faint(true).Italic(true).MarginLeft(2)
)

func (m ModifiersImpl) String() string {
	if len(m.AsJSON()) == 0 {
		return modEmptyStyle.Render("    No Modifiers") + "\n"
	}

	var lines []string
	for _, mods := range m.modifiers {
		for _, mod := range mods {
			lines = append(lines, "    "+mod.String())
		}
	}
	return strings.Join(lines, "\n") + "\n"
}


func NewModifiers() ModifiersImpl {
	return ModifiersImpl{
		modifiers: make(map[ModifierTarget][]Modifier),
	}
}
