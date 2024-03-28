package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

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

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

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
	fmt.Println(fibonacci(30))
}
