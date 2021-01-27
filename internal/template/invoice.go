package template

type InvoiceContainer struct {
	Invoice Invoice `yaml:"invoice"`
}

type Invoice struct {
	Billing  Billing  `yaml:"billing"`
	Client   Company  `yaml:"client"`
	Company  Company  `yaml:"company"`
	Contact  Contact  `yaml:"contact"`
	Metadata Metadata `yaml:"metadata"`
	Notes    string   `yaml:"notes"`
	Order    Order    `yaml:"order"`
}

type Address struct {
	Street   string `yaml:"street"`
	Area     string `yaml:"area"`
	PostCode string `yaml:"post-code"`
	Country  string `yaml:"coutnry"`
}

type Order struct {
	Lines     OrderLines `yaml:"lines"`
	Subtotal  string     `yaml:"subtotal"`
	TaxName   string     `yaml:"tax-name"`
	TaxRate   string     `yaml:"tax-rate"`
	TaxAmount string     `yaml:"tax-amount"`
	Total     string     `yaml:"total"`
}

type OrderLines []OrderLine

type OrderLine struct {
	Item     string `yaml:"item"`
	Quantity int    `yaml:"quantity"`
	Cost     string `yaml:"cost"`
	Total    string `yaml:"total"`
}

type Price struct {
	Value    int    `yaml:"value"`
	Currency string `yaml:"currency"`
}

type Contact struct {
	Email string `yaml:"email"`
	Phone string `yaml:"phone"`
}

type Company struct {
	Name    string  `yaml:"name"`
	TaxID   string  `yaml:"tax-id"`
	Address Address `yaml:"address"`
}

type Billing struct {
	IBAN string `yaml:"iban"`
	BIC  string `yaml:"bic"`
}

type Metadata struct {
	Number    int    `yaml:"number"`
	IssueDate string `yaml:"issue-date"`
	PayByDate string `yaml:"pay-by-date"`
}

type Tax struct {
	Name string  `yaml:"name"`
	Rate float64 `yaml:"rate"`
}
