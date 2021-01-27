package config

import (
	"errors"
	"testing"
	"time"

	"github.com/simonsargeant/fed/internal/template"
	"github.com/stretchr/testify/assert"
)

func TestData_Create(t *testing.T) {
	for s, tc := range map[string]struct {
		data   Data
		n      int
		t      time.Time
		client string
		items  map[string]int

		file string
		res  template.Invoice
		err  error
	}{
		"Doesn't create invoice when client invalid": {
			data: Data{
				Clients:   map[string]Client{"hedley-le-maistre": {}},
				Tax:       map[string]Tax{"lots": {}},
				Currency:  "EUR",
				Inventory: Items{"potato": {}},
			},
			n:      1,
			t:      time.Now(),
			client: "lily-langtree",
			items:  map[string]int{"potato": 3},

			err: errors.New("client \"lily-langtree\" not found"),
		},
		"Doesn't create invoice when tax invalid": {
			data: Data{
				Clients:   map[string]Client{"hedley-le-maistre": {Tax: "less"}},
				Tax:       map[string]Tax{"lots": {}},
				Currency:  "EUR",
				Inventory: Items{"potato": {}},
			},
			n:      1,
			t:      time.Now(),
			client: "hedley-le-maistre",
			items:  map[string]int{"potato": 3},

			err: errors.New("tax \"less\" for client \"hedley-le-maistre\" not found"),
		},
		"Doesn't create invoice when order invalid": {
			data: Data{
				Clients:   map[string]Client{"hedley-le-maistre": {Tax: "lots"}},
				Tax:       map[string]Tax{"lots": {}},
				Currency:  "EUR",
				Inventory: Items{"potato": {}},
			},
			n:      1,
			t:      time.Now(),
			client: "hedley-le-maistre",
			items:  map[string]int{"milk": 3},

			err: errors.New("create order for \"hedley-le-maistre\": item \"milk\" not found"),
		},
		"Creates invoice": {
			data: Data{
				Clients: map[string]Client{
					"hedley-le-maistre": {
						Company: Company{
							Name:  "Hedley Le Maistre",
							TaxID: "HLM123123",
							Address: Address{
								Street:   "123 Rue de la Corbiere",
								Area:     "St. Aubin",
								PostCode: "HE3 123",
								Country:  "Herm",
							},
						},
						InvoicePrefix: "HedleyLeMaistre",
						Tax:           "lots",
					},
				},
				Tax:      map[string]Tax{"lots": {Name: "Lots", Rate: 50}},
				Currency: "EUR",
				Inventory: Items{
					"potato": {
						Name:  "Potato",
						Price: 401,
					},
					"milk": {
						Name:  "Milk",
						Price: 3100,
					},
				},
				Billing: Billing{
					IBAN: "LH12 1234 1234 1234",
					BIC:  "LHBANK",
				},
				Company: Company{
					Name:  "Lihou Farms",
					TaxID: "LIF123123",
					Address: Address{
						Street:   "321 Rue de L'Etacq",
						Area:     "Gorey",
						PostCode: "LI1 321",
						Country:  "Lihou",
					},
				},
				Contact: Contact{
					Email: "lihou.farms@puffinmail.com",
					Phone: "+44123456789",
				},
			},
			n:      3,
			t:      time.Date(2020, 2, 22, 2, 2, 2, 2, time.UTC),
			client: "hedley-le-maistre",
			items:  map[string]int{"milk": 2, "potato": 3},

			file: "HedleyLeMaistre0003",
			res: template.Invoice{
				Billing: template.Billing{
					IBAN: "LH12 1234 1234 1234",
					BIC:  "LHBANK",
				},
				Client: template.Company{
					Name:  "Hedley Le Maistre",
					TaxID: "HLM123123",
					Address: template.Address{
						Street:   "123 Rue de la Corbiere",
						Area:     "St. Aubin",
						PostCode: "HE3 123",
						Country:  "Herm",
					},
				},
				Company: template.Company{
					Name:  "Lihou Farms",
					TaxID: "LIF123123",
					Address: template.Address{
						Street:   "321 Rue de L'Etacq",
						Area:     "Gorey",
						PostCode: "LI1 321",
						Country:  "Lihou",
					},
				},
				Contact: template.Contact{
					Email: "lihou.farms@puffinmail.com",
					Phone: "+44123456789",
				},
				Metadata: template.Metadata{
					Number:    3,
					IssueDate: "22/02/2020",
					PayByDate: "22/03/2020",
				},
				Notes: "This invoice refers to the period 01/02/2020 to 28/02/2020",
				Order: template.Order{
					Lines: []template.OrderLine{
						{
							Item:     "Milk",
							Quantity: 2,
							Cost:     "€31.00",
							Total:    "€62.00",
						},
						{
							Item:     "Potato",
							Quantity: 3,
							Cost:     "€4.01",
							Total:    "€12.03",
						},
					},
					Subtotal:  "€74.03",
					TaxName:   "Lots",
					TaxRate:   "50.0%",
					TaxAmount: "€37.02",
					Total:     "€111.05",
				},
			},
		},
	} {
		tc := tc
		t.Run(s, func(t *testing.T) {
			t.Parallel()

			file, res, err := tc.data.Create(tc.n, tc.t, tc.client, tc.items)

			if tc.err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.err.Error())
			}

			assert.Equal(t, tc.file, file)
			assert.Equal(t, tc.res, res)
		})
	}

}
