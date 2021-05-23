package main

import (
	"fmt"
	"os"
)

func Select(c chan int, quits chan struct{}){
	// select case if for channel and is similar to switch case
	for {
		select {
		case <- c:
			fmt.Println("received on channel c")
		case <- quits:
			fmt.Println("received on channel quits")
			os.Exit(1)
		}
	}
}
func selectExample(){

	c := make(chan int)
	quits := make(chan struct{})
	go func() {
		c <- 1
		quits <- struct{}{}
	}()
	go Select(c, quits)

	select {}
}