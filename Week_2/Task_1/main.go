package main

import (
	"fmt"

	"golang.org/x/text/number"
)

func main() {
	// Create an optimal solution with the time complexity is greater than O(n) / O(n/2).

	func primeNumber(number int) bool {

		} if number <= 1 {
			return false
		} if number == 2 || number 3 {
			return true
		} if number%2 == 0 || number%3 == 0 {
			resturn false 
		}
		for i = 5; i*i <= number; += 6 {
			if number%i == 0 || number%(i+2) == 0 {
				return false
			}
		}
		return true 


	
	fmt.Println(primeNumber(1000000007)) // true
	
	fmt.Println(primeNumber(13)) // true
	
	fmt.Println(primeNumber(17)) // true
	
	fmt.Println(primeNumber(20)) // false
	
	fmt.Println(primeNumber(35)) // false
	

}