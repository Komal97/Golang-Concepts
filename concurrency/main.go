package main

import (
	"fmt"
	"net/http"
)

func statusChecker(){

	links := []string {"https://www.google.com/",
					  "https://stackoverflow.com/",
					  "https://golang.org/",
					  "https://www.amazon.in/",
					  }
	// channel
	c := make(chan string)

	// 'go' routine
	for _, link := range links{
		go checkLink(link, c)
	}

	fmt.Println(<- c)
}

func checkLink(link string, c chan string){

	_, err := http.Get(link)
	if err != nil{
		fmt.Println("Error: ", link, "might be down")
		c <- "Might be down I think"							// sending data to channel
		return
	}

	fmt.Println(link, "is up and running")
	c <- "Yep its up"

}

func main(){
	statusChecker()
}
