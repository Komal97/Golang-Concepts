package main

import (
	"fmt"
	"net/http"
)


/*
Description
-----------
Make a GET call to every one time link and receive some data on channel in Go Routine.
Print received data on channel in main routine.

OUTPUT
------
https://www.google.com/ is up and running
Yep its up
https://stackoverflow.com/ is up and running
Yep its up
https://www.amazon.in/ is up and running
Yep its up
https://golang.org/ is up and running
Yep its up
*/

func statusChecker1(){

	links := []string {"https://www.google.com/",
		"https://stackoverflow.com/",
		"https://golang.org/",
		"https://www.amazon.in/",
	}
	// unbuffered channel = here channel can hold max 1 value at a time
	c := make(chan string)

	// 'go' routine
	for _, link := range links{
		go checkLink1(link, c)
	}

	//fmt.Println(<- c)          O/P = https://www.google.com/
	//fmt.Println(<- c)			 O/P = https://www.google.com/
	//fmt.Println(<- c)          O/P = https://golang.org/
	//fmt.Println(<- c)			 O/P = https://www.amazon.in/

	for i := 0; i < len(links); i++{
		fmt.Println(<- c)
	}
}

func checkLink1(link string, c chan string){

	_, err := http.Get(link)
	if err != nil{
		fmt.Println("Error: ", link, "might be down")
		c <- "Might be down I think"							// sending data to channel
		return
	}

	fmt.Println(link, "is up and running")
	c <- "Yep its up"

}
