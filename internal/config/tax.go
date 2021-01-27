package config

import (
	"fmt"

	"github.com/simonsargeant/fed/internal/template"
)

type Tax struct {
	Name string  `yaml:"name"`
	Rate float64 `yaml:"rate"`
}

func (t Tax) ToTemplate() template.Tax {
	return template.Tax{
		Name: t.Name,
		Rate: t.Rate,
	}
}

func (t Tax) ToString() string {
	return fmt.Sprintf("%s %.1f%%", t.Name, t.Rate)
}
