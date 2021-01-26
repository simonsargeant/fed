package config

import "github.com/Rhymond/go-money"

type Price int64

func (p Price) ToString(code string) string {
	return money.New(int64(p), code).Display()
}
