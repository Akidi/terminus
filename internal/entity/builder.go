// import path: terminus/internal/entity
// file path: ./internal/entity/builder.go
package entity

import (
	"terminus/internal/stats"

	"github.com/rs/xid"
)

type EntityBuilder interface {
	SetName(name string) EntityBuilder
	SetLevel(level int) EntityBuilder
	SetArchType(archtype ArchType) EntityBuilder
	SetAttributes(ac stats.Attributes) EntityBuilder
	Build() Entity
	Reset() EntityBuilder
}

type EntityBuilderImpl struct {
	entity *EntityImpl
}

func (eb *EntityBuilderImpl) SetName(name string) EntityBuilder {
	eb.entity.name = name
	return eb
}

func (eb *EntityBuilderImpl) SetLevel(level int) EntityBuilder {
	eb.entity.level = level
	return eb
}

func (eb *EntityBuilderImpl) SetArchType(archType ArchType) EntityBuilder {
	eb.entity.archType = archType
	return eb
}

func (eb *EntityBuilderImpl) SetAttributes(ac stats.Attributes) EntityBuilder {
	eb.entity.attributes = ac
	return eb
}

func (eb *EntityBuilderImpl) Reset() EntityBuilder {
	eb.entity = &EntityImpl{}
	return eb
}

func (eb *EntityBuilderImpl) Build() Entity {
	result := eb.entity
	result.id = xid.New().String()
	if result.name == "" {
		result.name = "Unnamed"
	}
	if result.level == 0 {
		result.level = 1
	}
	if result.archType == "" {
		result.archType = Adventurer
	}
	if result.attributes == nil {
		result.attributes = stats.NewAttributes()
	}
	eb.Reset()
	return result
}

func NewEntityBuilder() EntityBuilder {
	return &EntityBuilderImpl{entity: &EntityImpl{}}
}