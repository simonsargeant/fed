package config

import (
	"fmt"
	"time"

	"github.com/simonsargeant/fed/internal/template"
)

type Data struct {
	Billing   Billing           `yaml:"billing"`
	Inventory Items             `yaml:"inventory"`
	Clients   map[string]Client `yaml:"clients"`
	Company   Company           `yaml:"company"`
	Contact   Contact           `yaml:"contact"`
	Currency  string            `yaml:"currency"`
	Tax       map[string]Tax    `yaml:"tax"`
}

func (d Data) Create(n int, t time.Time, clientName string, items map[string]int) (string, template.Invoice, error) {
	client, ok := d.Clients[clientName]
	if !ok {
		return "", template.Invoice{}, fmt.Errorf("client %q not found", clientName)
	}

	tax, ok := d.Tax[client.Tax]
	if !ok {
		return "", template.Invoice{}, fmt.Errorf("tax %q for client %q not found", client.Tax, clientName)
	}

	lines, err := d.Inventory.CreateOrder(items, d.Currency, tax)
	if err != nil {
		return "", template.Invoice{}, fmt.Errorf("create order for %q: %s", clientName, err)
	}

	start, end := MonthPeriod(t)

	return fmt.Sprintf("%s%04d", client.InvoicePrefix, n), template.Invoice{
		Metadata: template.Metadata{
			Number:    n,
			IssueDate: printDate(t),
			PayByDate: printDate(t.AddDate(0, 1, 0)),
		},
		Billing: d.Billing.ToTemplate(),
		Client:  client.Company.ToTemplate(),
		Company: d.Company.ToTemplate(),
		Contact: d.Contact.ToTemplate(),
		Order:   lines,
		Notes:   fmt.Sprintf("This invoice refers to the period %s to %s", printDate(start), printDate(end)),
	}, nil
}
