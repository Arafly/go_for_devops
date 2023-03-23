// The bufio package also makes available primitives that are used to scan and tokenize buffered input data from an io.Reader source. 
// The bufio.Scanner type scans input data using the Split method to define tokenization strategies.

// This time, the code uses bufio.Scanner (instead of the fmt.Fscan function) to scan the content of the text file using the bufio.ScanLines function:

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	fmt.Printf(
		"%-10s %-10s %-6s %-6s\n",
		"Planet", "Diameter", "Moons", "Ring?",
	)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		fmt.Printf(
			"%-10s %-10s %-6s %-6s\n",
			fields[0], fields[1], fields[2], fields[3],
		)
	}
}

// Using bufio.Scanner is done in four steps as shown in the previous example:

// First, use bufio.NewScanner(io.Reader) to create a scanner
// Call the scanner.Split method to configure how the content is tokenized
// Traverse the generated tokens with the scanner.Scan method
// Read the tokenized data with the scanner.Text method
