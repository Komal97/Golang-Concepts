package main

import (
	"fmt"
)

func main() {

	// variable declaration
	// method - 1
	// var a string = "a"

	// method - 2
	a := "a"
	fmt.Println(a)

	// reassignment
	a = "b"
	fmt.Println(a)

	// function call
	a = test()
	fmt.Println(a)

	fmt.Print("--------------- Check ODD EVEN ---------------")
	var nums []int
	for i:= 0; i < 10; i++{
		nums = append(nums, i)
	}

	for _, num := range nums {
		if num % 2 == 0 {
			fmt.Printf("\n%v is even", num)
		} else{
			fmt.Printf("\n%v is odd", num)
		}
	}

	implementArrayAndSlice()
	implementMap()
	implementStruct()
}

func test() string {
	return "c"
}
