package main

import (
	"fmt"
	"math"
)

func main() {

	const Pi float64 = math.Pi // Pi constant from the math package

	// Variables (you can modify these values)
	var radius float64 = 2.0 // radius of the tube
	var height float64 = 5.0 // height of the tube

	// Calculate the volume of the tube
	volume := Pi * radius * radius * height

	// Print the result
	fmt.Printf("Volume of the tube: %.2f cubic units\n", volume)

}
