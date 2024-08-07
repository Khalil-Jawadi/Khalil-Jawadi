package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(catalan(7))  // 429
	fmt.Println(catalan(10)) // 16796
}

func catalan(n int) *big.Int {
	// Calculating (2n)!
	twoNFactorial := factorial(2 * n)

	// Calculating (n+1)!
	nPlusOneFactorial := factorial(n + 1)

	// Calculating n!
	nFactorial := factorial(n)

	// Catalan number = (2n)! / ((n+1)! * n!)
	result := new(big.Int).Mul(nPlusOneFactorial, nFactorial)
	result.Div(twoNFactorial, result)

	return result
}

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
