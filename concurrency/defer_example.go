package main

import "fmt"

/*
OUTPUT
------
doing some stuff

// defer statements run in LIFO manner
Are we Done?
Done!
*/

func deferExample(){

	// defer statement delays the execution until nearby (deferExample) function return
	defer fmt.Println("Done!")
	defer fmt.Println("Are we Done?")
	fmt.Println("doing some stuff")
}
