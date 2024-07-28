package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  // [[katak, civic], kasur, rusak]
	fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) // [[racecar, civic, kak], seru, kasur, bilik]
	fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  // [[civic, hahah], masuk, garam]
}

func groupPalindrome(words []string) []any {
	isPalindrome := func(word string) bool {
		runes := []rune(word)
		n := len(runes)
		for i := 0; i < n/2; i++ {
			if runes[i] != runes[n-1-i] {
				return false
			}
		}
		return true
	}

	palindromes := []string{}
	nonPalindromes := []string{}

	for _, word := range words {
		if isPalindrome(word) {
			palindromes = append(palindromes, word)
		} else {
			nonPalindromes = append(nonPalindromes, word)
		}
	}

	result := []any{}
	if len(palindromes) > 0 {
		result = append(result, palindromes)
	}
	for _, word := range nonPalindromes {
		result = append(result, word)
	}

	return result
}
