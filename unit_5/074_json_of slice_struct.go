package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type location struct {
		Name string  `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	locations := []location{
		{Name: "curiosity", Lat: 32.32, Long: -43.32},
		{Name: "discovery", Lat: 142.32, Long: -153.32},
		{Name: "opportunity", Lat: 120.21, Long: 214.32},
	}

	for i := range locations {
		// I think for this particular example it's ok to disregard error handling
		//bytes, err := json.Marshal(locations[i])
		//error_handling(err)

		bytes, _ := json.Marshal(locations[i])
		fmt.Println(string(bytes))
	}
}

//func error_handling(err error) {
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//}
