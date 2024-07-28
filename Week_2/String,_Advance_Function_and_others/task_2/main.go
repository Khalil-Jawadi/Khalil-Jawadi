package main

import (
	"fmt"
)

func main() {
	fmt.Println(spinString("alta"))    // aalt
	fmt.Println(spinString("alterra")) // aalrtre
	fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(word string) string {
	runes := []rune(word)
	length := len(runes)
	result := make([]rune, length)

	for i, r := range runes {
		newIndex := (2 * i) % length
		result[newIndex] = r
	}

	return string(result)
}
