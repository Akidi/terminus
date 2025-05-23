// import path: terminus/internal/modifier
// file path: ./internal/modifiers/director.go
package modifiers

type Director struct {
	builder ModifierBuilder
}

func NewDirector(builder ModifierBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) BuildModifier(name string, target ModifierTarget, kind ModifierKind, value float64, source ModifierSourceType) Modifier {
	return d.builder.
		Reset().
		SetName(name).
		SetTarget(target).
		SetKind(kind).
		SetValue(value).
		SetSource(source).
		Build()
}