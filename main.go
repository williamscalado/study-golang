package main

import (
	"fmt"
	"studygo/calc"
	"studygo/helpers"
)

var mock = calc.Coordenadas{
	Lat:  44545,
	Long: 45454,
}

var valuesFormat = helpers.PropsMoneyFormat{
	NumberValue: "32.50",
	Currency:    "EUR",
}

var valueMoney = helpers.MoneyFormat(valuesFormat)

func main() {

	fmt.Println(valueMoney)
}
