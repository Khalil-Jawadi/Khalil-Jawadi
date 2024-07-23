package main

import "fmt"

func main() {

	var score int
	fmt.Println("enter your score")
	fmt.Scan(&score)

	if score >= 85 && score <= 100 {
		fmt.Println("Grade A")
	} else if score >= 70 && score <= 84 {
		fmt.Println("Grade B")
	} else if score >= 69 && score <= 55 {
		fmt.Println("Grade C")
	} else if score >= 40 && score <= 50 {
		fmt.Println("Grade D")
	} else if score >= 0 && score <= 39 {
		fmt.Println("Grade E")
	} else if score >= 100 && score <= 0 {
		fmt.Println("Invalid score")
	}

}
