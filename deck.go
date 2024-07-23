package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

//shuffle logic
// for each index. card in cards
// generate a randmo number between 0 and len(cards)-1
// swap the current card and the card at cards[randomNumber]

func (d deck) shuffleDeck() {
	// establish a new rand type with differnet source
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// fancy one liner to swap elements
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
