package config

import "github.com/simonsargeant/fed/internal/template"

type Billing struct {
	IBAN string `yaml:"iban"`
	BIC  string `yaml:"bic"`
}

func (b Billing) ToTemplate() template.Billing {
	return template.Billing{
		IBAN: b.IBAN,
		BIC:  b.BIC,
	}
}
