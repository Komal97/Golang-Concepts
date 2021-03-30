/*
Code to test the deck
 */
package main

import (
	"os"
	"testing"
)

// function to test newDeck() function
func TestNewDeck(t *testing.T){

	d := newDeck()

	// Case 1: Check if deck contains 'x' cards or not
	if len(d) != 16{
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	// Case 2: Check first card is "Ace of Spades" or not
	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	// Case 3: Check last card is "Four of Clubs" or not
	if d[len(d)-1] != "Four of Clubs"{
		t.Errorf("Expected last card of Four of Clubs but got %v", d[len(d)-1])
	}
}

// function to test saveToFile() and readFromFile() functions
func TestSaveToFileAndReadFromFile(t *testing.T){

	// clean and remove files
	os.Remove("_decktesting")

	// Case 1: Save deck to file
	saveDeck := newDeck()
	saveDeck.saveToFile("_decktesting")

	// Case 2: Load deck from file
	loadedDeck := readFromFile("_decktesting")
	if len(loadedDeck) != 16{
		t.Errorf("Expected deck length of 16 but got %v", len(loadedDeck))
	}

	// clean and remove files
	os.Remove("_decktesting")
}

