package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Quote struct represents a quote with ID, text, and author
type Quote struct {
	ID     int
	Text   string
	Author string
	Movie  string
}

// global variable to keep track of the last assigned ID
var lastQuoteID int

// PopulateQuotes populates a slice of Quote structs from an array of text and authors
func PopulateQuotes(text, authors, movies []string) ([]Quote, error) {
	var quotes []Quote

	// Check if the length of text and authors arrays are equal
	if len(text) != len(authors) || len(text) != len(movies) {
		return nil, errors.New("Error: Length of text, authors and movies arrays must be equal")
	}

	// Iterate over the text and authors arrays to create Quote structs
	for i := 0; i < len(text); i++ {
		quote := Quote{
			ID:     lastQuoteID + 1,
			Text:   text[i],
			Author: authors[i],
			Movie:  movies[i],
		}
		quotes = append(quotes, quote)
		lastQuoteID++
	}

	return quotes, nil
}

func getRandomQuoteIndex(numQuotes int) (int, error) {
	if numQuotes == 0 {
		return -1, errors.New("Error: No quotes available")
	}
	// Generate a random number between 0 and numQuotes
	randomIndex := rand.Intn(numQuotes)

	return randomIndex, nil
}

func getData() ([]string, []string, []string, error) {
	var textData, authorsData, moviesData []string

	file, err := os.Open("quotes.txt")
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// each line is in the format text|author|movie
		parts := strings.Split(line, "|")
		if len(parts) == 3 {
			textData = append(textData, parts[0])
			authorsData = append(authorsData, parts[1])
			moviesData = append(moviesData, parts[2])
		}
	}

	err = scanner.Err()
	if err != nil {
		return nil, nil, nil, err
	}

	return textData, authorsData, moviesData, nil

}

func main() {
	// Reset lastQuoteID to 0 at the beginning of the program
	lastQuoteID = 0

	// Get data from the file
	text, authors, movies, err := getData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Populate quotes from existing text and authors
	quotes, err := PopulateQuotes(text, authors, movies)
	if err != nil {
		fmt.Println("Error:", err)
		return
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
	fmt.Printf("%s -%s, %s\n", randomQuote.Text, randomQuote.Author, randomQuote.Movie)
}
