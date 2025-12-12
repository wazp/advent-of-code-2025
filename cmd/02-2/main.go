package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// file to read line by line from a text file and process each line
var filePath = "cmd/02-1/input.txt"
var current = 0

func isRepeatedPattern(num int) bool {
	s := strconv.Itoa(num)
	length := len(s)

	for size := 1; size <= length/2; size++ {
		if length%size != 0 {
			continue
		}
		part := s[:size]
		ok := true
		for i := size; i < length; i += size {
			if s[i:i+size] != part {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func processRecord(line string) {
	parts := strings.Split(line, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	for i := start; i <= end; i++ {
		if isRepeatedPattern(i) {
			current += i
		}
	}
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
		line := scanner.Text()
		items := strings.Split(line, ",")
		for _, item := range items {
			processRecord(item)
		}
	}

	fmt.Printf("\nFinal Current: %d\n", current)

	//fmt.Printf("Final Current: %d, Number of times at 0: %d\n", current, numOfZero)

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during scanning: %v", err)
	}
}
