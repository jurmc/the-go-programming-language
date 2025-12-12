package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("AbsoluteZeroC: %g, AbsoluteZeroF: %g\n", AbsoluteZeroC, CToF(AbsoluteZeroC))
	fmt.Printf("FreezingC: %g, FreezingF: %g\n", FreezingC, CToF(FreezingC))
	fmt.Printf("BoilingC: %g, BoilingF: %g\n", BoilingC, CToF(BoilingC))

	fmt.Println(AbsoluteZeroC)
	fmt.Println(FreezingC)
	fmt.Println(BoilingC)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32) * 5 / 9)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g stopni C", c)
}
