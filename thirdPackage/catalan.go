package main

import (
	"fmt"
	"math/big"
)

func catalan(n int64) *big.Int {
	one := big.NewInt(1)

	denominator := factorial(2 * n)
	divisor := one.Mul(factorial(n+1), factorial(n))
	return one.Div(denominator, divisor)
}

func factorial(n int64) *big.Int {
	sum := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		sum.Mul(sum, big.NewInt(i))
	}
	return sum
}

func main() {
	for i := int64(0); i < 20; i++ {
		fmt.Println(catalan(i))
	}
}
