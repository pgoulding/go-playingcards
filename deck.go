package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

//d is reference to 'self'

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// _ specifiies an unused variable
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// deal(d deck, handSize)
// d = this deck
// deck = type of argument
// handSize = passed in argument of int type
// (deck, deck)
// returns two values of type deck
func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		// opt #1 - log error and return a call to the new deck
		fmt.Println("Error: ", err)
		// return newDeck()
		// opt #2 - log the error and quit the program
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
