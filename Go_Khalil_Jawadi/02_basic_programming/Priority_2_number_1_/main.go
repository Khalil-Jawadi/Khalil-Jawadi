package main

import "fmt"

func main() {

	var number int
	fmt.Println("enter a number: ")
	fmt.Scan(&number)

	// find the number factors
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			//check the factor
			// if the factor is even ..> print I
			// otherwise print O
			if i%2 == 0 {
				fmt.Print("I")
			} else {
				fmt.Print("O")
			}

		}

	}

}
