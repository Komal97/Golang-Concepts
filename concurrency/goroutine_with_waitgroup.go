package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanUp(){
	if r := recover(); r != nil{
		fmt.Println("Recovered in cleanup:", r)
	}
}

func say(s string){
	// Dpne() tells wait group that current go routine is finished
	defer wg.Done()
	defer cleanUp()
	for i := 0; i < 3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if  i == 2 {
			panic("Oops, i = 2 hits")
		}
	}
}

func waitGroupExample(){

	// adding go routine to waitgroup so that both can run concurrently
	wg.Add(2)
	//wg.Add(1)
	go say("hey")
	//wg.Add(1)
	go say("there")

	// it keeps on waiting until above 2 go routines are done
	wg.Wait()

	// this will run after above 2 go routines are finished
	wg.Add(1)
	go say("Hi")
	wg.Wait()
}


