package main

import (
	"fmt"
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
}
