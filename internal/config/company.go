package config

import "github.com/simonsargeant/fed/internal/template"

type Company struct {
	Name    string  `yaml:"name"`
	TaxID   string  `yaml:"tax-id"`
	Address Address `yaml:"address"`
}

func (c Company) ToTemplate() template.Company {
	return template.Company{
		Name:    c.Name,
		TaxID:   c.TaxID,
		Address: c.Address.ToTemplate(),
	}
}
