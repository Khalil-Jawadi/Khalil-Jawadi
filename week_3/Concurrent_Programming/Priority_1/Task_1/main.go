package main

import (
	"fmt"
	"time"
)

func reverseWord(word string) {
	for i := 1; i <= len(word); i++ {
		go func(substring string) {
			time.Sleep(3 * time.Second)
			fmt.Println(substring)
		}(word[:i])
	}
	time.Sleep(time.Duration(len(word)) * 3 * time.Second) // Wait for all goroutines to complete
}

func main() {
	word := "phone"
	reversed := ""
	for _, char := range word {
		reversed = string(char) + reversed
	}
	reverseWord(reversed)
}
