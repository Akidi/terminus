// import path: terminus/internal/modifier
// file path: ./internal/modifiers/modifier.go
package modifiers

type ModifierTarget string

const (
	TargetStrength			ModifierTarget = "strength"
	TargetDexterity			ModifierTarget = "dexterity"
	TargetVitality			ModifierTarget = "vitality"
	TargetIntelligence	ModifierTarget = "intelligence"
	TargetWisdom				ModifierTarget = "wisdom"
	TargetLuck					ModifierTarget = "luck"

	TargetMaxHP					ModifierTarget = "max_hp"
	TargetMaxMP					ModifierTarget = "max_mp"
)

type ModifierSourceType int

const (
	SourceItem		ModifierSourceType = iota
	SourceSkill
	SourceBuff
	SourcePassive
)

type ModifierKind int

const (
	Flat ModifierKind = iota
	Percent
	Multiplier
)

type Modifier interface {
	ID() 			string
	Name()		string
	Target()	ModifierTarget
	Kind()		ModifierKind
	Value()		float64
	Source()	ModifierSourceType
	Apply(base float64) float64
	AsJSON()	ModifierJson
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
	return base
}

func (m *ModifierImpl) AsJSON() ModifierJson {
	return ModifierJson{
		ID: m.id,
		Name: m.name,
		Target: m.target,
		Kind: m.kind,
		Value: m.value,
		Source: m.source,
	}
}

type ModifierJson struct {
	ID     string		`json:"id"`
	Name   string		`json:"name"`
	Target ModifierTarget	`json:"target"`
	Kind   ModifierKind	`json:"kind"`
	Value  float64	`json:"value"`
	Source ModifierSourceType	`json:"source"`
}
