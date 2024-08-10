package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Function to count word frequencies concurrently
func countWords(words <-chan string, frequencies chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	wordCount := make(map[string]int)
	for word := range words {
		wordCount[word]++
	}
	frequencies <- wordCount
}

// Function to write word frequencies to a CSV file
func writeCSV(filename string, wordCount map[string]int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for word, count := range wordCount {
		record := []string{word, fmt.Sprintf("%d", count)}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Create channels and wait group
	words := make(chan string, 100)
	frequencies := make(chan map[string]int)
	var wg sync.WaitGroup

	// Start goroutine to count words
	wg.Add(1)
	go countWords(words, frequencies, &wg)

	// Open and read the text document
	file, err := os.Open("sample.txt") // Change to your file path
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Fields(line) {
			words <- strings.ToLower(word) // Convert to lowercase for case-insensitive counting
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Close the words channel after reading the file
	close(words)

	// Wait for the word counting to finish and retrieve the result
	go func() {
		wg.Wait()
		close(frequencies)
	}()

	wordCount := make(map[string]int)
	for wc := range frequencies {
		for word, count := range wc {
			wordCount[word] += count
		}
	}

	// Write the results to a CSV file
	if err := writeCSV("word_frequencies.csv", wordCount); err != nil {
		fmt.Println("Error writing CSV file:", err)
	} else {
		fmt.Println("Word frequencies have been written to word_frequencies.csv")
	}
}
