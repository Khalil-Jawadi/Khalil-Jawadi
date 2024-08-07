package main

import "fmt"

func main() {
	fmt.Println(getSequence(4))   // 10
	fmt.Println(getSequence(15))  // 120
	fmt.Println(getSequence(100)) // 5050
}

func getSequence(n int) int {
	//0 1 3 6 10
	// +1  +2 +3 +4
	result := 0
	for i := 0; i <= n; i++ {
		result += i

	}

	return result

}
