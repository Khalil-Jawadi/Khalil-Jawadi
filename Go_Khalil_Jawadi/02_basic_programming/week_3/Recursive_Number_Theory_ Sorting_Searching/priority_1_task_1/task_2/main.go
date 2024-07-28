package main

import "fmt"

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
}

func getInfo(students []Student) {
	var bestMathStudent, bestScienceStudent, bestEnglishStudent, bestAverageStudent Student
	var highestMath, highestScience, highestEnglish, highestAverage float64

	for _, student := range students {
		if student.MathScore > highestMath {
			highestMath = float64(student.MathScore)
			bestMathStudent = student
		}

		if student.ScienceScore > highestScience {
			highestScience = float64(student.ScienceScore)
			bestScienceStudent = student
		}

		if student.EnglishScore > highestEnglish {
			highestEnglish = float64(student.EnglishScore)
			bestEnglishStudent = student
		}

		averageScore := (float64(student.MathScore) + float64(student.ScienceScore) + float64(student.EnglishScore)) / 3
		if averageScore > highestAverage {
			highestAverage = averageScore
			bestAverageStudent = student
		}
	}

	fmt.Printf("best student in math: %s with %.0f\n", bestMathStudent.Name, highestMath)
	fmt.Printf("best student in science: %s with %.0f\n", bestScienceStudent.Name, highestScience)
	fmt.Printf("best student in english: %s with %.0f\n", bestEnglishStudent.Name, highestEnglish)
	fmt.Printf("best student in average: %s with %.0f\n", bestAverageStudent.Name, highestAverage)
}
