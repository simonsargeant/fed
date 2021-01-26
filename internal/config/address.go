package config

import "github.com/simonsargeant/fed/internal/template"

type Address struct {
	Street   string `yaml:"street"`
	Area     string `yaml:"area"`
	PostCode string `yaml:"post-code"`
	Country  string `yaml:"country"`
}

func (a Address) ToTemplate() template.Address {
	return template.Address{
		Street:   a.Street,
		Area:     a.Area,
		PostCode: a.PostCode,
		Country:  a.Country,
	}
}
