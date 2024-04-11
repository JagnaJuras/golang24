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

// Obliczanie silnii
func factorial(n int64) *big.Int {
	fact := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		fact.Mul(fact, big.NewInt(i))
	}
	return fact
}

// Funkcja sprawdzająca, czy liczba znajduje się w dłuższym ciągu znaków (silnii)
func strIndex(s, substr string) int {
	n := len(substr)
	for i := 0; i <= len(s)-n; i++ {
		if s[i:i+n] == substr {
			return i
		}
	}
	return -1
}

// Funkcja sprawdzająca, czy wszystkie liczby z listy znajdują się w silnii
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

// Funkcja szukająca silnii zawierającej wszystkie liczby z listy
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

// Zainicjowanie mapy dla l. wywołań funkcji fibonacci dla poszczególnych liczb
var fibCalls map[int]int

func init() {
	fibCalls = make(map[int]int)
}

// Obliczanie n-tej liczby Fibonacciego
func fibonacci(n int) int {
	fibCalls[n]++

	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Wyświetlanie mapy z l. wywołań
func printFibCalls() {
	var ns []int
	for n := range fibCalls {
		ns = append(ns, n)
	}
	sort.Ints(ns)

	fmt.Println("Liczba wywołań funkcji fibonacci dla każdego n:")
	for _, n := range ns {
		fmt.Printf("n = %d: %d\n", n, fibCalls[n])
	}
}

// Funkcja do wyszukiwania liczby zawierającej l. wywołań najbliższą danej liczbie (Silnej Liczbie)
func findClosestFibCalls(n int64) int {
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
	// Utworzenie nicku
	var name, surname string
	fmt.Print("Imię: ")
	fmt.Scanln(&name)
	fmt.Print("Nazwisko: ")
	fmt.Scanln(&surname)
	nick := strings.ToLower(name[0:3] + surname[0:3])
	fmt.Println("Nick:", nick)

	// Stworzenie listy z kodami ASCII liter z nicku
	runes := []rune(nick)
	fmt.Println("Kody numeryczne ASCII:", runes)

	// Stworzenie listy z kodami ASCII w formie stringów
	var numStrings []string
	for _, r := range runes {
		numStrings = append(numStrings, strconv.Itoa(int(r)))
	}
	n := findFactorialContainingNumbers(numStrings)
	fmt.Printf("Twoja Silna Liczba: %d\n", n)
	fibonacci(30)
	// printFibCalls()
	closestFib := findClosestFibCalls(n)
	fmt.Printf("Twoja Słaba Liczba: %d\n", closestFib)
}
