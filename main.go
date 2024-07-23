package main

// newDeck -> Creates and Returns a list of playing cards
// print -> log the contents of a deck of cards
// shuffle -> shuffles all the cards in a deck
// deal -> create a hand of cards
// saveToFile -> save a list of cards to a file on the local machine
// newDeckFromFile -> Load a list of cards from the local machine.

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}
