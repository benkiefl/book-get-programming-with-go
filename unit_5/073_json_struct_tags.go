package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type location struct {
		// requires CamelCase so fields can be exported for json package
		Lat, Long float64
	}

	// you might want snake_case esp when working in conjunction with ruby/python
	// you can manually alter the output via struct tags

	type location_2 struct {
		// struct tags are formatted as key:"value"
		// key tends to be the name of a package
		// you could even customize the Lat field for both JSON and XML by doing
		// `json:"latitude" xml:"latitude"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
		// raw string literals (``) are best, because you don't need escape sequences
		// otherwise it would be "json:\"latitude\""
	}

	curiosity := location{-3.23, 153.32}
	discovery := location_2{-3.23, 153.32}

	// func json.Marshal(v interface{}) ([]byte, error)
	// Marshal returns the JSON encoding of v.
	bytes, err := json.Marshal(curiosity)
	bytes_2, err_2 := json.Marshal(discovery)

	exitOnError(err)
	exitOnError(err_2)

	fmt.Println(string(bytes))
	fmt.Println(string(bytes_2))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
