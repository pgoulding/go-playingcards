package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	//create new deck
	d := newDeck()
	// write an if statement to see if the deck has the right number of cards
	if len(d) != 52 {
		t.Errorf("Expected length of 52 but got: %v", len(d))
		// if fails, tell the test handler something went wrong
	}

	//write an if statement to see if the first card in deck is correct
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// write an if statement to check last element in deck
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got: %v", d[len(d)-1])
	}

}

// ridiculously long name, but hey, it's descriptive
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	// write an if statement for testing the creation of deck file
	//clean up any created testing files, ii
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("expected 52 cards in deck , got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
