package main

import "fmt"

func SimpleEquations(a, b, c int) {
	for x := 1; x <= 100; x++ {
		for y := 1; y <= 100; y++ {
			for z := 1; z <= 100; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Printf("%d %d %d\n", x, y, z)
					return
				}
			}
		}
	}
	fmt.Println("No solution.")
}

func main() {
	SimpleEquations(1, 2, 3)  // No solution.
	SimpleEquations(6, 6, 14) // 1 2 3
}
