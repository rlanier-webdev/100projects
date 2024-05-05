**README**

This is a simple Go program that generates random quotes and serves them through an HTTP server. It reads quote data from a file named "quotes.txt" and populates a slice of Quote structs, each representing a quote with an ID, text, author, and movie.

### Features:
- **Random Quote Generation:** The program selects a random quote from the populated list of quotes and displays it.
- **HTTP Server:** Quotes are served through an HTTP server running on port 8000. When accessing the root URL ("/"), a random quote is displayed.
- **File Input:** Quote data is read from a file named "quotes.txt", with each line containing a quote in the format: text|author|movie.

### How to Use:
1. Ensure you have Go installed on your machine.
2. Create a file named "quotes.txt" and populate it with quotes in the format mentioned above.
3. Run the program by executing the command `go run main.go`.
4. Open a web browser and navigate to `http://localhost:8000` to view a random quote.

### Files:
- **main.go:** Contains the main code for the program.
- **quotes.txt:** Contains the quote data in the required format.

### Dependencies:
- This program does not have any external dependencies beyond the standard Go libraries.

### Notes:
- Ensure that the "quotes.txt" file exists and is correctly formatted before running the program.
- The program will log any errors encountered during execution.