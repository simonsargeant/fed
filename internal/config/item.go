package config

import (
	"fmt"

	"github.com/simonsargeant/fed/internal/template"
)

type Items map[string]Item

type Item struct {
	Name  string `yaml:"name"`
	Price Price  `yaml:"price"`
}

func (i Items) CreateOrder(items map[string]int, currency string, tax Tax) (template.Order, error) {
	var lines []template.OrderLine

	var subtotal Price
	var taxAmount Price
	var total Price
	for name, quantity := range items {
		item, ok := i[name]
		if !ok {
			return template.Order{}, fmt.Errorf("item %q not found", name)
		}

		lineTotal := Price(quantity) * item.Price
		subtotal += lineTotal

		taxAmount := Price(0.0)
		if tax.Rate != 0.0 {
			taxAmount := Price(float32(lineTotal) * tax.Rate / 100.0)
			lineTotal += taxAmount
		}

		taxAmount += taxAmount
		total += lineTotal

		lines = append(lines, template.OrderLine{
			Item:     item.Name,
			Quantity: quantity,
			Cost:     item.Price.ToString(currency),
			Total:    lineTotal.ToString(currency),
		})

	}

	order := template.Order{
		Lines:     lines,
		Subtotal:  subtotal.ToString(currency),
		TaxName:   tax.Name,
		TaxRate:   fmt.Sprintf("%.1f", tax.Rate) + "%",
		TaxAmount: taxAmount.ToString(currency),
		Total:     total.ToString(currency),
	}

	return order, nil
}
