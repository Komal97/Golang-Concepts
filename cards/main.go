/*
Code to create and manipulate deck
*/

package main

import "fmt"

func main() {

	//cards := []string{"Ace of diamonds", newCard()}
	//cards := deck{"Ace of diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	//
	//for index, card := range cards {
	//	fmt.Println(index, card)
	//}

	cards := newDeck()
	//fmt.Println("---------- Using Receiver Function -----------")
	//cards.printCards()
	//cards.print()

	hand, remCards := cards.deal(5)
	fmt.Println("---------- Hand -----------")
	hand.printCards()
	fmt.Println("---------- Remaining cards -----------")
	remCards.printCards()

	fmt.Println("---------- Save and Read from File -----------")
	cards.saveToFile("cards.txt")
	cards = readFromFile("cards.txt")
	cards.printCards()

	fmt.Println("---------- Shuffle cards -----------")
	cards.shuffle()
	cards.printCards()

}

