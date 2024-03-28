package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

// ===================== SILNA LICZBA =====================

func factorial(n int64) *big.Int {
	fact := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		fact.Mul(fact, big.NewInt(i))
	}
	return fact
}

func containsAllNumbers(factorialStr string, numbers []string) bool {
	for _, num := range numbers {
		index := strIndex(factorialStr, num)
		if index == -1 {
			return false
		}
		factorialStr = factorialStr[:index] + factorialStr[index+len(num):]
	}
	return true
}

func strIndex(s, substr string) int {
	n := len(substr)
	for i := 0; i <= len(s)-n; i++ {
		if s[i:i+n] == substr {
			return i
		}
	}
	return -1
}

func findFactorialContainingNumbers(numbers []string) int64 {
	var num int64
	for {
		fact := factorial(num)
		factStr := fact.String()
		if containsAllNumbers(factStr, numbers) {
			break
		}
		num++
	}
	return num
}

// ===================== SŁABA LICZBA =====================

var fibCalls map[int]int

func init() {
	fibCalls = make(map[int]int)
}

func fibonacci(n int) int {
	// Increment the counter for this value of n
	fibCalls[n]++

	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func printFibCalls() {
	fibonacci(30)

	// Sort the n values
	var ns []int
	for n := range fibCalls {
		ns = append(ns, n)
	}
	sort.Ints(ns)

	// Print the counts
	fmt.Println("Number of calls for each n:")
	for _, n := range ns {
		fmt.Printf("n = %d: %d calls\n", n, fibCalls[n])
	}
}

func findClosestFibonacci(n int64) int {
	closestN := 0
	closestDifference := math.MaxInt64

	for fibN, count := range fibCalls {
		difference := int(math.Abs(float64(count) - float64(n)))
		if difference < closestDifference {
			closestDifference = difference
			closestN = fibN
		}
	}

	return closestN
}

// ===================== DZIAŁANIE =====================

func main() {
	name := "Jagna"
	surname := "Juras"
	nick := strings.ToLower(name[0:3] + surname[0:3])
	fmt.Println("Nickname:", nick)

	runes := []rune(nick)
	fmt.Println("Runes:", runes)

	var numStrings []string
	for _, r := range runes {
		numStrings = append(numStrings, strconv.Itoa(int(r)))
	}
	n := findFactorialContainingNumbers(numStrings)
	fmt.Printf("The smallest factorial number containing all numbers %v is %d\n", numStrings, n)
	fibonacci(30)
	closestFib := findClosestFibonacci(n)
	fmt.Printf("The number in Fibonacci sequence closest to the result is: %d\n", closestFib)
}
