package main

import (
	"fmt"
	"net/http"
	"time"
)


/*
Description
-----------
Make a GET call to every link and once routine corresponding to that link is finished, it calls itself again and again

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

This output keeps on printing infinite times
*/

// commented code is method - 2 for repeating routines
func statusChecker2(){

	links := []string {"https://www.google.com/",
		"https://stackoverflow.com/",
		"https://golang.org/",
		"https://www.amazon.in/",
	}
	// channel
	c := make(chan string)

	// created go routine once for each link
	for _, link := range links{
		go checkLink2(link, c)
	}

	// run an inifite loop
	// link 47 is blocked because here it is waiting for data on channel
	// as soon as data received, then a new go routine with same link is again generated and previous one is finished.

	// method - 1
	//for {
	//	go checkLink2(<- c, c)
	//}
	for l := range c {
		go func(newLink string) {
			// sleep pauses the current go routine for atleast duration d
			time.Sleep(5 * time.Second)
			checkLink2(newLink, c)
		}(l)
		//time.Sleep(5 * time.Second)
		//go checkLink2(l, c)
	}

	// method - 2 (using infite loop inside checklink function)
	//for {
	//	fmt.Println(<- c)
	//}
}

func checkLink2(link string, c chan string){
	//for {
		_, err := http.Get(link)
		if err != nil {
			fmt.Println("Error: ", link, "might be down")
			// c <- "Might be down I think"
			c <- link // sending link to channel
			return
		}

		fmt.Println(link, "is up and running")
		// c <- "Yep its up"
		c <- link // sending link to channel
	//}
}
