package main

import "fmt"

func implementMap()  {
	fmt.Println("\n--------------- MAP -----------------")

	// CREATE
	// method - 1
	colors1 := map[string] string{
		"red": "#ff0000",
		"green": "#4bf745",
	}
	printMap(colors1)

	// method - 2
	//var colors2 map[string] string
	//colors2["red"] = "#ff0000"
	//colors2["green"] = "#4bf745"
	//fmt.Println(colors2)

	// method - 3
	colors3 := make(map[string] string)
	colors3["red"] = "#ff0000"
	colors3["green"] = "#4bf745"
	colors3["white"] = "#ffffff"
	printMap(colors3)

	// DELETE KEY
	delete(colors3, "green")
	printMap(colors3)

}

func printMap(c map[string]string){

	for color, hex := range c{
		fmt.Println("Hex code for", color, "is", hex)
	}
}