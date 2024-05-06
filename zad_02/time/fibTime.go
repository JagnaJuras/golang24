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

// Funkcja Ackermanna
func ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	} else if n == 0 {
		return ackermann(m-1, 1)
	} else {
		return ackermann(m-1, ackermann(m, n-1))
	}
}

func measureTime(m, n int) time.Duration {
	start := time.Now()
	ackermann(m, n)
	return time.Since(start)
}

func main() {
	// numbers := []int{40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	numbers := []int{}

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

	testCases := [][]int{
		{3, 5},
		{3, 6},
		{3, 7},
		{3, 8},
		{3, 9},
		{3, 10},
		{3, 11},
		{3, 12},
		{3, 13},
	}

	for _, testCase := range testCases {
		m, n := testCase[0], testCase[1]
		duration := measureTime(m, n)
		fmt.Printf("Czas obliczania funkcji ackermann(%d, %d): %v\n", m, n, duration)
	}
}
