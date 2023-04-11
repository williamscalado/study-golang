package helpers

import (
	"github.com/bojanz/currency"
)

type PropsMoneyFormat struct {
	NumberValue string
	Currency    string
}

func MoneyFormat(props PropsMoneyFormat) string {
	if props.Currency == "" {
		props.Currency = "BRL"
	}
	locale := currency.NewLocale("tr")
	formatter := currency.NewFormatter(locale)
	amount, _ := currency.NewAmount(props.NumberValue, props.Currency)
	return formatter.Format(amount)
}
