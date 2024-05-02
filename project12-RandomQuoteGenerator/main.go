package main

import (
	"fmt"
	"math/rand"
)

// Quote struct represents a quote with ID, text, and author
type Quote struct {
	ID     int
	Text   string
	Author string
}

// global variable to keep track of the last assigned ID
var lastQuoteID int

// PopulateQuotes populates a slice of Quote structs from an array of text and authors
func PopulateQuotes(text, authors []string) []Quote {
	var quotes []Quote

	// Check if the length of text and authors arrays are equal
	if len(text) != len(authors) {
		fmt.Println("Error: Length of text and authors arrays must be equal")
		return nil
	}

	// Iterate over the text and authors arrays to create Quote structs
	for i := 0; i < len(text); i++ {
		quote := Quote{
			ID:     lastQuoteID + 1,
			Text:   text[i],
			Author: authors[i],
		}
		quotes = append(quotes, quote)
		lastQuoteID++
	}

	return quotes
}

func getRandomQuoteIndex(numQuotes int) (int, error) {
	if numQuotes == 0 {
		return -1, fmt.Errorf("Error: No quotes available")
	}
	// Generate a random number between 0 and numQuotes
	randomIndex := rand.Intn(numQuotes)

	return randomIndex, nil
}

func main() {
	// Reset lastQuoteID to 0 at the beginning of the program
	lastQuoteID = 0

	// Existing quotes
	text := []string{"Ray", "Joy"}
	authors := []string{"This is a test.", "This is another test."}

	// Populate quotes from existing text and authors
	quotes := PopulateQuotes(text, authors)

	// Print the quotes w/IDs
	fmt.Println("Quotes:")
	for _, quote := range quotes {
		fmt.Printf("ID: %d, Text: %s, Author: %s\n", quote.ID, quote.Text, quote.Author)
	}

	// Get the number of quotes
	numQuotes := len(quotes)

	// Generate a random quote index
	randomIndex, err := getRandomQuoteIndex(numQuotes)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Access the random quote using the index
	randomQuote := quotes[randomIndex]

	fmt.Println("\nRandom Quote:")
	fmt.Printf("ID: %d, Text: %s, Author: %s\n", randomQuote.ID, randomQuote.Text, randomQuote.Author)
}
