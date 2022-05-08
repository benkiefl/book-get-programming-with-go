package main

import (
	"fmt"
)

type kelvin float64
type celsius float64
type fahrenheit float64

// 0: kelvin to celsius
// 1: celsius to kelvin
// 2: kelvin to fahrenheit
// 3: fahrenheit to kelvin
// 4: celsius to fahrenheit
// 5: fahrenheit to celsius
const choice int64 = 0

// k_to_c method converts kelvin to celsius
func (k kelvin) k_to_c() celsius {
	return celsius(k - 273.15)
}

// c_to_k method converts celsius to kelvin
func (c celsius) c_to_k() kelvin {
	return kelvin(c + 273.15)
}

// k_to_f method converts kelvin to fahrenheit
func (k kelvin) k_to_f() fahrenheit {
	return fahrenheit((k-273.15)*1.8 + 32)
}

// f_to_k method converts fahrenheit to kelvin
func (f fahrenheit) f_to_k() kelvin {
	return kelvin(((f - 32) / 1.8) + 273.15)
}

// c_to_f method converts celsius to fahrenheit
func (c celsius) c_to_f() fahrenheit {
	return fahrenheit((c * 1.8) + 32)
}

// f_to_c method converts fahrenheit to celsius
func (f fahrenheit) f_to_c() celsius {
	return celsius((f - 32) / 1.8)
}

// draw_table calls functions to draw table
func draw_table() {
	draw_header()
	draw_conversion()
	draw_separator()
}

// draw_header draws table header
func draw_header() {
	draw_separator()
	switch choice {
	case 0:
		fmt.Printf("| K\t| °C\t\t|\n")
	case 1:
		fmt.Printf("| °C\t| K\t\t|\n")
	case 2:
		fmt.Printf("| K\t| °F\t\t|\n")
	case 3:
		fmt.Printf("| °F\t| K\t\t|\n")
	case 4:
		fmt.Printf("| °C\t| °F\t\t|\n")
	case 5:
		fmt.Printf("| °F\t| °C\t\t|\n")
	default:
		fmt.Printf("NON-DEFINED VALUE FOR VARIABLE CHOICE\n")
	}
	draw_separator()
}

// draw_k_to_c draws conversion table
// conversion based on value of const CHOICE
func draw_conversion() {
	switch choice {
	case 0:
		for i := 0; i <= 400; i += 20 {
			k := kelvin(i)
			fmt.Printf("| %4.0f\t| %6.2f\t|\n", k, k.k_to_c())
		}
	case 1:
		for i := -40; i <= 100; i += 5 {
			c := celsius(i)
			fmt.Printf("| %4.0f\t| %6.2f\t|\n", c, c.c_to_k())
		}
	case 2:
		for i := 0; i <= 400; i += 20 {
			k := kelvin(i)
			fmt.Printf("| %4.0f\t| %6.2f\t|\n", k, k.k_to_f())
		}
	case 3:
		for i := -40; i <= 100; i += 5 {
			f := fahrenheit(i)
			fmt.Printf("| %4.0f\t| %6.2f\t|\n", f, f.f_to_k())
		}
	case 4:
		for i := -40; i <= 100; i += 5 {
			c := celsius(i)
			fmt.Printf("| %4.0f\t| %6.2f\t|\n", c, c.c_to_f())
		}
	case 5:
		for i := -40; i <= 100; i += 5 {
			f := fahrenheit(i)
			fmt.Printf("| %4.0f\t| %6.2f\t|\n", f, f.f_to_c())
		}
	case 6:
	case 7:
	default:
		fmt.Printf("NON-DEFINED VALUE FOR VARIABLE CHOICE\n")
	}
}

// draw_separator draws simple ===== line of appropriate length
func draw_separator() {
	// str := "="
	// fmt.Println(strings.Repeat(str, 25))
	fmt.Printf("=========================\n")
}

// main function
func main() {

	draw_table()
}
