// import path: terminus/internal/skills
// file path: ./internal/skills/skills.go
package skills

import "encoding/json"

type SkillName int

const (
	OneHanded SkillName = iota
	TwoHanded
	Staves
	Ranged
	Magic

	Barter
)

type Skill interface {
	Name() string
	ToJSON() ([]byte, error)
}

type SkillImpl struct {
	name string
}

func (s *SkillImpl) Name() string {
	return s.name
}

type skillJSON struct {
	Name string `json:"name"`
}

func (s *SkillImpl) ToJSON() ([]byte, error) {
	v := skillJSON{
			Name: s.name,
	}
	return json.Marshal(v)
}