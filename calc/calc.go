package calc

import "math"

type iMath struct {
	Pi     float32
	Sqrt2  float32
	SqrtPi float32
}

func Sum(num1, num2 uint) uint {
	return num1 + num2
}

func MathVariables() iMath {
	return iMath{Pi: math.Pi, Sqrt2: math.Sqrt2, SqrtPi: math.SqrtPi}
}

const pi = math.Pi

func Multiply(num1, num2 float32) float32 {
	return num1 * num2 * pi
}

type Coordenadas struct {
	Lat  int
	Long int
}
