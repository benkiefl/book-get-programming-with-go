package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64

// it's actually quite simple after a while, a method is
// a function with the return (receiver) stated before the name
// and the method being accessed via dot notation
// + a method requires a newly declared type for it to work
// that's all

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(((k - 273.15) * 1.8) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 1.8) + 32.0)
}

func (f fahrenheit) kelvin() kelvin {
	return kelvin(((f - 32.0) / 1.8) + 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) / 1.8)
}

func main() {

	var k kelvin = 100
	var c celsius = 100
	var f fahrenheit = 100

	fmt.Println("kelvin", k, "to celsius", k.celsius(), "to", k.fahrenheit())
	fmt.Println("celsius", c, "to kelvin", c.kelvin(), "to fahrenheit", c.fahrenheit())
	fmt.Println("fahrenheit", f, "to kelvin", f.kelvin(), "to celsius", f.celsius())

}
