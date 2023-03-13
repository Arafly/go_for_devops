package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Flag to find csv, then Parse the csv file
func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	// timer for the quiz
	timeLimit := flag.Int("limit", 30, "time limit for quiz in sec")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		dismiss(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFile))
	}

	// Create a reader to read the file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		dismiss("Could not parse the provided CSV file")
	}
	Question := parseLines(lines)
	// fmt.Println(Question)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// Keep track of the score
	correct := 0

problemLoop:
	// Looping through the questions and verify the answer
	for q, a := range Question {
		fmt.Printf("Problem #%d: %s = \n", q+1, a.question)
		// Create a channel to receive the answer
		answerCh := make(chan string)
		// write a go routine to read the answer
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
			// Waiting for a message on the time channel. code will block here until it gets a message from the channel
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			// Checking if the answer is correct
			if answer == a.answer {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Incorrect!")
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(Question))
}


func parseLines(lines [][]string) []Question {
	ret := make([]Question, len(lines))
	for i, line := range lines {
		ret[i] = Question{
			question: line[0],
			// Remove the whitespace in the answer
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type Question struct {
	question string
	answer string
}

// Exit function
func dismiss(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}