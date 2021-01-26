package config

type Client struct {
	InvoicePrefix string  `yaml:"invoice-prefix"`
	Company       Company `yaml:"company"`
	Tax           string  `yaml:"tax"`
}
