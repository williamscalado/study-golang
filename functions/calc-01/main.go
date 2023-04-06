package main

import (
	"fmt"
	"log"
)

func sumValues(number1, number2 uint) uint {
	if number1 < 0 || number2 < 0 {
		log.Print("Valor não aceito")
		return 0
	}

	return number1 + number2
}

func main() {
	result := sumValues(30, 60)
	result1 := sumValues(20, 25)
	fmt.Println(result + result1)
}
