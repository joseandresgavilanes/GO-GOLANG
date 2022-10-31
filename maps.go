package main

import "fmt"

func main() {
	//MAPS

	//Declaration
	var map_1 = map[string]int32{
		"Car":      50000,
		"House":    20000,
		"Computer": 1000,
	}

	fmt.Println(map_1)

	//Another type of declaration
	map_2 := make(map[int]int)
	map_2[10] = 1
	map_2[12] = 4
	map_2[8] = 5
	map_2[11] = 2

	//Displaying a map
	for i, v := range map_2 {
		fmt.Println(i, "=", v) //There is no order at the momento of displaying a map
	}

	//Knowing if a value exist
	//If we display map_1["Car"] it will show the value
	value := map_1["Car"]
	fmt.Println(value)

	//But what if it doesn't exist
	value = map_1["aasd"]
	fmt.Println(value) //It will show zero

	value, ok := map_1["aasd"] //If we want to know if it exists, it must be done with the second returing value
	//it will return a bool depending on if exists or not
	fmt.Println(value, ok)
}