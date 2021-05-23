package main

import "sync"

//import (
//	"fmt"
//	"sync"
//)

var wg2 sync.WaitGroup

func foo(c chan int, value int){
	defer wg2.Done()
	c <- value * 5
}

//// buffered channel = here channel can hold more than 1 value at a time
//func main(){
//	fooChan := make(chan int, 10)  // add buffer space for 10 items
//
//	for i := 0; i < 10; i++{
//		wg2.Add(1)
//		go foo(fooChan, i)
//	}
//
//	wg2.Wait()
//	close(fooChan)
//
//	for item := range fooChan{
//		fmt.Println(item)
//	}
//}