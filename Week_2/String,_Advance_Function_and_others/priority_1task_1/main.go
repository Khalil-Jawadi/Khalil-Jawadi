package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pickVocals("alterra academy"))                    // aea aae
	fmt.Println(pickVocals("i love programming"))                 // i oe oai
	fmt.Println(pickVocals("go is awesome programming language")) // o i aeoe oai auae
}

func pickVocals(sentence string) string {
	vowels := "aeiouAEIOU"
	isVowel := func(r rune) bool {
		return strings.ContainsRune(vowels, r)
	}

	words := strings.Fields(sentence)
	var result []string

	for _, word := range words {
		var vowelOnlyWord strings.Builder
		for _, char := range word {
			if isVowel(char) {
				vowelOnlyWord.WriteRune(char)
			}
		}
		result = append(result, vowelOnlyWord.String())
	}

	return strings.Join(result, " ")
}
