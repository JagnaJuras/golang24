package main

import (
	"fmt"
	"math/big"
	"time"
)

// Obliczanie n-tej liczby Fibonacciego rekurencyjnie
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	numbers := []int{40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, num := range numbers {
		start := time.Now()
		fibonacci(num)
		elapsed := time.Since(start)
		fmt.Printf("Czas obliczania funkcji fibonacci(%d): %s\n", num, elapsed)
	}
	baseTime := big.NewInt(58) // w sekundach
	targetN := 258

	estimatedTime := new(big.Int).Mul(baseTime, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(targetN-50)), nil))
	fmt.Printf("Szacowany czas obliczania fibonacci(258): %v seconds\n", estimatedTime)
	estimatedTimeInYears := new(big.Rat).SetFrac(estimatedTime, big.NewInt(86400*365.25)) // 86400 sekund * 365.25 dni = 1 rok

	fmt.Printf("Szacowany czas obliczania fibonacci(258): %v years\n", estimatedTimeInYears.FloatString(2))
}
