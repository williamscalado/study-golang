package main

import (
	"fmt"
	"studygo/calc"
)

var mock = calc.Coordenadas{
	Lat:  44545,
	Long: 45454,
}

func main() {
	fmt.Println(calc.Sum(10, 10))
	fmt.Println(calc.Multiply(10, 20))
	fmt.Println(calc.Sub(10, 20))
	var result *calc.Coordenadas = &mock
	fmt.Println(result.Lat)
	data := calc.MathVariables()
	fmt.Println(data.Sqrt2)
}
