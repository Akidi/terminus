// import path: terminus/internal/entity
// file path: ./internal/entity/entity.go
package entity

import (
	"terminus/internal/common"
	"terminus/internal/stats"
)


type Entity interface {
	ID() string
	Name() string
	Level() int
	ArchType() ArchType
	Attributes() stats.Attributes
	ToJSON() ([]byte, error)
}

type EntityImpl struct {
	id					string
	name				string
	archType		ArchType
	level				int
	attributes	stats.Attributes
}

func (e *EntityImpl) ID() string {
	return e.id
}

func (e *EntityImpl) Name() string {
	return e.name
}

func (e *EntityImpl) Level() int {
	return e.level
}

func (e *EntityImpl) ArchType() ArchType {
	return e.archType
}

func (e *EntityImpl) Attributes() stats.Attributes {
	return e.attributes
}

type EntityJson struct {
	ID					string 		`json:"id"`
	Name				string		`json:"name"`
	ArchType		ArchType	`json:"archtype"`
	Level				int				`json:"level"`
	Attributes	[]stats.AttributeJson `json:"attributes"`
}

func (e *EntityImpl) ToJSON() ([]byte, error) {
	return common.ToJSON(e)
}

func (e *EntityImpl) AsJSON() *EntityJson {
	return &EntityJson{
		ID: e.id,
		Name: e.name,
		ArchType: e.archType,
		Level: e.level,
		Attributes: e.attributes.AsJSON(),
	}
}