package main

import (
	"fmt"
	"os"
)

// newDeck -> Creates and Returns a list of playing cards
// shuffle -> shuffles all the cards in a deck
// deal -> create a hand of cards
// saveToFile -> save a list of cards to a file on the local machine
// print -> log the contents of a deck of cards
// newDeckFromFile -> Load a list of cards from the local machine.

func main() {
	//cards := newDeckFromFile("my_cards")
	cards := newDeck()
	// cards.shuffleDeck()
	// cards.print()
	// fmt.Println(len(cards))

	for {
		// print out a menu of choices
		fmt.Println("\n--- Playing Cards Menu ---")
		fmt.Println("1. Create a new deck")
		fmt.Println("2. Shuffle the deck")
		fmt.Println("3. Deal a hand")
		fmt.Println("4. Save deck to file")
		fmt.Println("5. Show Remaining Cards")
		fmt.Println("6. Quit")
		fmt.Print("Enter your choice: ")

		//establish choice var based on input
		var choice int
		_, err := fmt.Scanf("%d", &choice)

		// error handling
		if err != nil {
			fmt.Println("Invalid input, please enter a number between 1 and 6.")
			continue
		}

		// swtich statewment for choice
		switch choice {
		case 1:
			fmt.Println("Creating new deck")
			cards = newDeck()
			fmt.Printf("Created deck with a length of %v cards", len(cards))
		case 2:
			fmt.Println("Shuffling the deck...")
			cards.shuffleDeck()
			fmt.Println("Deck has been shuffled")
		case 3:
			fmt.Println("Deal a hand...")
			if len(cards) >= 5 {
				hand, newCards := deal(cards, 5)
				cards = newCards
				fmt.Println(hand)
			} else {
				fmt.Println("Not engouh cards left to deal, try creating a new deck...")
			}

		case 4:
			fmt.Println("Save deck to file...")
			cards.saveToFile("_deckfile")
		case 5:
			fmt.Println("Showing remaining cards in the deck...")
			fmt.Printf("There are %v cards left in deck", len(cards))
			cards.print()
		case 6:
			fmt.Println("Quitting the program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please enter a number between 1 and 6.")
		}

	}
}
