package main

import "fmt"

type celsius float64

// method fahrenheit converts from celsius to fahrenheit
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

type fahrenheit float64

// method celsius converts from fahrenheit to celsius
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n" // that line is interesting
	numberFormat = "%.1f"
)

// creating your own function type is helpful for passing functions around
// makes code more readable, but it's not strictly necessary
// type getRowFn func(row int) (string, string)

// drawTable draws a two column table
// w/ above commented out type defintion we could just do:
// 		func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
// but if we only use it once, I don't see a need, I would rather do:
func drawTable(hdr1, hdr2 string, rows int, getRow func(int) (string, string)) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2) // awesome that works!
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func ctof(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}

func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func main() {
	drawTable("°C", "°F", 29, ctof)
	fmt.Println()
	drawTable("°F", "°C", 29, ftoc)
}
