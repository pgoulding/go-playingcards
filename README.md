# Playing Cards Project

This project provides a Go implementation for managing and manipulating a deck of playing cards. It includes functionalities for creating a deck, shuffling, drawing cards, and more.

## Features

- Create a standard 52-card deck
- Shuffle the deck
- Save the deck to file
- Draw cards from the deck
- Reset the deck

## Installation

To use this project, you need to have Go installed on your machine. You can download and install Go from the [official website](https://golang.org/dl/).

Once Go is installed, you can clone the repository and install the dependencies:

```bash
git clone https://github.com/pgoulding/go-playingcards
cd go-playingcards
go mod tidy
```

## Running the Project

You can run the project using the go run command:

```bash
go run main.go
```

## Testing

This project includes unit tests for the various functionalities. To run the tests, use the go test command:

```bash
go test ./...
```

## Usage

Here are some examples of how to use the different functionalities in this project:

## Menu

Upon starting the program you should be shown a list of Options that yuou can select by entering a numerical menu option:

```bash
--- Playing Cards Menu ---
1. Create a new deck
2. Shuffle the deck
3. Deal a hand
4. Save deck to file
5. Show Remaining Cards
6. Quit
```

### 1. Create a new deck

This creates a brand new unshuffled deck.

### 2. Shuffle the deck

This takes the exsiting deck and shuffles the cards in a random order.

### 3. Deal a hand

This takes out 5 cards from the deck and shows them to you.

### 4. Save deck to file

This saves the existing deck to a "_deckfile" file.

### 5. Show remaining cards

This shows the remaining cards in the deck.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.