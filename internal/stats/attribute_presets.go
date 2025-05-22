// import path: terminus/internal/stats
// file path: ./internal/stats/attribute_presets.go
package stats

func NewWarriorAttributes() Attributes {
	return &AttributesImpl{attributes: map[AttributeName]Attribute{
		"Strength":     NewAttributeFromDef(AttributeDefs["Strength"], 8),
		"Vitality":     NewAttributeFromDef(AttributeDefs["Vitality"], 5),
		"Dexterity":    NewAttributeFromDef(AttributeDefs["Dexterity"], 4),
		"Luck":         NewAttributeFromDef(AttributeDefs["Luck"], 2),
		"Wisdom":       NewAttributeFromDef(AttributeDefs["Wisdom"], 1),
		"Intelligence": NewAttributeFromDef(AttributeDefs["Intelligence"], 1),
	}}
}

func NewMageAttributes() Attributes {
	return &AttributesImpl{attributes: map[AttributeName]Attribute{
		"Intelligence": NewAttributeFromDef(AttributeDefs["Intelligence"], 8),
		"Wisdom":       NewAttributeFromDef(AttributeDefs["Wisdom"], 6),
		"Luck":         NewAttributeFromDef(AttributeDefs["Luck"], 4),
		"Dexterity":    NewAttributeFromDef(AttributeDefs["Dexterity"], 2),
		"Strength":     NewAttributeFromDef(AttributeDefs["Strength"], 1),
		"Vitality":     NewAttributeFromDef(AttributeDefs["Vitality"], 1),
	}}
}

func NewBattleMageAttributes() Attributes {
	return &AttributesImpl{attributes: map[AttributeName]Attribute{
		"Intelligence": NewAttributeFromDef(AttributeDefs["Intelligence"], 6),
		"Strength":     NewAttributeFromDef(AttributeDefs["Strength"], 6),
		"Wisdom":       NewAttributeFromDef(AttributeDefs["Wisdom"], 4),
		"Dexterity":    NewAttributeFromDef(AttributeDefs["Dexterity"], 3),
		"Vitality":     NewAttributeFromDef(AttributeDefs["Vitality"], 1),
		"Luck":         NewAttributeFromDef(AttributeDefs["Luck"], 1),
	}}
}

func NewRogueAttributes() Attributes {
	return &AttributesImpl{attributes: map[AttributeName]Attribute{
		"Dexterity":    NewAttributeFromDef(AttributeDefs["Dexterity"], 8),
		"Luck":         NewAttributeFromDef(AttributeDefs["Luck"], 6),
		"Strength":     NewAttributeFromDef(AttributeDefs["Strength"], 3),
		"Vitality":     NewAttributeFromDef(AttributeDefs["Vitality"], 2),
		"Wisdom":       NewAttributeFromDef(AttributeDefs["Wisdom"], 1),
		"Intelligence": NewAttributeFromDef(AttributeDefs["Intelligence"], 1),
	}}
}

func NewClericAttributes() Attributes {
	return &AttributesImpl{attributes: map[AttributeName]Attribute{
		"Wisdom":       NewAttributeFromDef(AttributeDefs["Wisdom"], 8),
		"Intelligence": NewAttributeFromDef(AttributeDefs["Intelligence"], 5),
		"Vitality":     NewAttributeFromDef(AttributeDefs["Vitality"], 4),
		"Strength":     NewAttributeFromDef(AttributeDefs["Strength"], 2),
		"Luck":         NewAttributeFromDef(AttributeDefs["Luck"], 1),
		"Dexterity":    NewAttributeFromDef(AttributeDefs["Dexterity"], 1),
	}}
}
