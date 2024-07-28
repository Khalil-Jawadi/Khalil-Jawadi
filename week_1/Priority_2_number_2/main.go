package main

import "fmt"

func main() {
	var budget, duration, difficulty int
	fmt.Println(" enter a budget: ")
	fmt.Scan(&budget)

	fmt.Printf("enter a duration: ")
	fmt.Scan(&duration)

	fmt.Println("project difficulty: ")
	fmt.Scan(&difficulty)

	var budgetScore, durationScore, difficultyScore int

	// TODO: assign the budget scorebudget
	if budget >= 0 && budget <= 50 {
		budgetScore = 4
	} else if budget >= 51 && budget <= 80 {
		budgetScore = 3
	} else if budget >= 81 && budget <= 100 {
		budgetScore = 2
	} else if budget > 100 {
		budgetScore = 1
	}

	//TODO: assign the duration score
	if duration >= 0 && duration <= 20 {
		durationScore = 10
	} else if duration >= 21 && duration <= 30 {
		durationScore = 5
	} else if duration >= 31 && duration <= 50 {
		durationScore = 2
	} else if duration > 50 {
		durationScore = 1
	}

	//TODO: assign the difficulty score
	if difficulty >= 0 && difficulty <= 3 {
		difficultyScore = 10
	} else if difficulty >= 4 && difficulty <= 6 {
		difficultyScore = 5
	} else if difficulty >= 8 && difficulty <= 10 {
		difficultyScore = 1
	} else if difficulty > 10 {
		difficultyScore = 0
	}

	var priorityScore int = budgetScore + durationScore + difficultyScore

	// TODO: check the priority of the project
	if priorityScore >= 17 && priorityScore <= 24 {
		fmt.Println("High")
	} else if priorityScore >= 10 && priorityScore <= 16 {
		fmt.Print("Medium")
	} else if priorityScore >= 3 && priorityScore <= 9 {
		fmt.Println("impossible")
	}

}
