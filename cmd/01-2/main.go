package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// file to read line by line from a text file and process each line
var filePath = "cmd/01-2/input.txt"
var current = 50
var passedZero = 0

func processLine(line string) {
	action := line[0]
	var value int
	fmt.Sscanf(line[1:], "%d", &value)

	for i := 0; i < value; i++ {
		if action == 'L' {
			current = (current + 1) % 100
		} else {
			current = (current - 1 + 100) % 100
		}

		if current == 0 {
			passedZero++
		}
	}

	// fmt.Printf("Action %c Value %d -> current %d (zero hits so far: %d)\n", action, value, current, passedZero)
}

func main() {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new Scanner for the filei
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		processLine(line)
	}

	fmt.Printf("Final Current: %d, Passed zero: %d\n", current, passedZero)

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during scanning: %v", err)
	}
}
