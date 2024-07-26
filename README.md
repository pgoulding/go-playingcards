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

### Creating a Deck

```go
deck := NewDeck()
```

### Shuffling the Deck

```go
deck.Shuffle()
```

### Drawing Cards

```go
card, err := deck.Draw()
if err != nil {
    log.Fatalf("Error drawing card: %v", err)
}
fmt.Println("Drew card:", card)
```

### Resetting the Deck

```go
deck.Reset()
```

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.