package config

import "github.com/simonsargeant/fed/internal/template"

type Contact struct {
	Email string `yaml:"email"`
	Phone string `yaml:"phone"`
}

func (c Contact) ToTemplate() template.Contact {
	return template.Contact{
		Email: c.Email,
		Phone: c.Phone,
	}
}
