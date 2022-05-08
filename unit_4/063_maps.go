package main

import "fmt"

func main() {

	// mean temperatures from
	// https://nssdc.gsfc.nasa.gov/planetary/factsheet/

	// map[key]value
	// you can always look up stuff using the keys
	planets_temp := map[string]int{
		"mercury": 167,
		"venus":   464,
		"earth":   15,
		"moon":    -20,
		"mars":    -65,
		"jupiter": -110,
		"saturn":  -140,
		"uranus":  -195,
		"neptune": -200,
		"pluto":   -225,
	}

	fmt.Println(planets_temp["mercury"])

	// access non-existing key yields the default zero value for type, here 0 (int)
	fmt.Println(planets_temp["sun"])

	// so how to differentiate between default value or real value of 0?
	// comma, ok syntax (don't have to call the variables that, but it's known under that idiom)
	if sun, ok := planets_temp["sun"]; ok {
		// sun will be value for key, if found
		// check will be true or false whether key exists or not
		fmt.Printf("Mean temperature of Sun is %v\n", sun)
	} else {
		fmt.Printf("there is no \"sun\" in temperature database\n")
	}

	// for visualization of things
	var_a, var_b := planets_temp["test"]
	fmt.Println(var_a, var_b)

	// maps aren't copied

	some_map := map[string]int{
		"A": 0,
		"B": 1,
	}

	some_map_2 := some_map

	some_map_2["B"] = 0

	fmt.Println(some_map, some_map_2)

	delete(some_map_2, "B")

	fmt.Println(some_map, some_map_2)

	// similar to how various slices can map to the same underlying array

	// unless initialized with a composite literal (similar to slices)
	// maps need to be allocated with make function
	// make() for maps takes two arguments, second preallocates for number of keys
	db_users := make(map[string]int, 2) // much like cap 10 for slices, length always 0

	fmt.Println(db_users)

	// something like cap(db_users) doesn' work, there's no equivalent
	// but maps grow dynamically as key-value pairs are added, so no need

	// so what's the need for pre-allocation to begin with?
	// imagine you read from a slice and transfer to a map
	// instead of allocating additional space every time you add a pair
	// you can pre-allocate (memory optimization)

	db := make(map[string]int, 5)
	s0 := []string{"a", "b", "c", "d", "e"}
	s1 := []int{0, 1, 2, 3, 4}

	for i, _ := range s0 {
		// you could just write for i := range s0 {
		db[s0[i]] = s1[i]
	}

	fmt.Println(db)

}
