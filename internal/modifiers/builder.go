// import path: terminus/internal/modifier
// file path: ./internal/modifiers/builder.go
package modifiers

import "github.com/rs/xid"

type ModifierBuilder interface {
	SetName(name string) ModifierBuilder
	SetTarget(target ModifierTarget) ModifierBuilder
	SetKind(kind ModifierKind) ModifierBuilder
	SetValue(value float64) ModifierBuilder
	SetSource(source ModifierSourceType) ModifierBuilder
	Reset() ModifierBuilder
	Build() Modifier
}

type ModifierBuilderImpl struct {
	modifier *ModifierImpl
}

func (mb *ModifierBuilderImpl) SetName(name string) ModifierBuilder {
	mb.modifier.name = name
	return mb
}

func (mb *ModifierBuilderImpl) SetTarget(target ModifierTarget) ModifierBuilder {
	mb.modifier.target = target
	return mb
}

func (mb *ModifierBuilderImpl) SetKind(kind ModifierKind) ModifierBuilder {
	mb.modifier.kind = kind
	return mb
}

func (mb *ModifierBuilderImpl) SetValue(value float64) ModifierBuilder {
	mb.modifier.value = value
	return mb
}

func (mb *ModifierBuilderImpl) SetSource(source ModifierSourceType) ModifierBuilder {
	mb.modifier.source = source
	return mb
}

func (mb *ModifierBuilderImpl) Reset() ModifierBuilder {
	mb.modifier = &ModifierImpl{}
	return mb
}

func (mb *ModifierBuilderImpl) Build() Modifier {
	result := *mb.modifier
	result.id = xid.New().String()
	mb.Reset()
	return &result
}

func NewModifier() ModifierBuilder {
	return &ModifierBuilderImpl{modifier: &ModifierImpl{}}
}

