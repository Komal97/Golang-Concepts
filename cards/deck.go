/*
Code to describe what the deck is and how it works
 */
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of 'deck' which is a slice of strings
type deck[] string

// function to create and return new deck
func newDeck() deck{

	cards := deck{}

	cardSuits := []string {"Spades", "Diamond", "Hearts", "Clubs"}
	cardsValues := []string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardsValues{
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

// receiver function to print cards
func (d deck) print(){
	fmt.Println(d)
}

func (d deck) printCards(){

	// (d deck) receives deck type variable (object) as parameter like "self" in python
	for i, card := range d{
		fmt.Println(i, card)
	}
	// d[0] = "ABC" 	// change value in actual object as deck is slice actually whose address is passed to params
}

// function to find subset of deck to play with and create 2 different sets (subset and remaining cards)
func (d deck) deal(handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

// convert deck to string
func (d deck) toString () string{

	// convert deck -> string[] -> string (comma separated)
	return strings.Join([]string (d), ",")

}

// function to save data to file and return error if exist
func (d deck) saveToFile(fileName string) error{

	// WriteFile (filename, [] byte, permission)
	return ioutil.WriteFile(fileName, []byte (d.toString()), 0666)
}

// function to read from file
func readFromFile(fileName string) deck {

	bs, err := ioutil.ReadFile(fileName)   // return byte slice and error
	if err != nil{
		// Option 1 -> log the error and return a call to newDeck()
		// Option 2 -> log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)                  // 0 means success
	}

	//  convert byte slice -> string -> [] string -> deck
	s := strings.Split(string (bs), ",")
	return deck (s)
}

// function to shuffle cards
func (d deck) shuffle(){


	// method - 1
	//source := rand.NewSource(time.Now().UnixNano())
	//r := rand.New(source)

	// method - 2
	rand.Seed(time.Now().UnixNano())
	max := len(d)
	for i := range d{
		//newPosition := r.Intn(max-i)+i
		newPosition := rand.Intn(max-i)+i		// generate random number between i and len(deck)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}