package main

import "fmt"

func implementArrayAndSlice(){

	// array
	fmt.Println("\n--------------- ARRAY -----------------")
	//ar := [4] string {"f", "g", "h", "i"}
	var ar [4] string
	ar[0] = "f"
	ar[1] = "g"
	ar[2] = "h"

	for i:= 0; i<len(ar); i++{
		fmt.Print(ar[i] + " ")
	}

	// slice
	fmt.Println("\n--------------- SLICE -----------------")
	//arr := []string{test(), "d"}
	arr := []string{"z", "d"}
	arr = append(arr, "e")     // return a new slice by appending new values in old slice
	fmt.Println(arr)

	for index, val := range arr {
		fmt.Println(index, val)
	}

	for i:= 0; i<len(arr); i++{
		fmt.Print(arr[i] + " ")
	}
}


