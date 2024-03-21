package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

func factorial(number int64) *big.Int {
	var fact = new(big.Int)
	fact.MulRange(1, number)
	return (fact)
}

func containsAllRunes(numStr string, runes []rune) bool {
	for _, r := range runes {
		if !strings.ContainsRune(numStr, r) {
			return false
		}
	}
	return true
}

func findFactorialWithRunes(runes []rune) int64 {
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	num := int64(1)
	for {
		fact := factorial(num)
		if containsAllRunes(fact.String(), runes) {
			return num
		}
		num++
	}
}

func main() {
	name := "Jagna"
	surname := "Juras"
	nick := strings.ToLower(name[0:3] + surname[0:3])
	fmt.Println(nick)
	runes := []rune(nick)
	for i := 0; i < len(runes); i++ {
		fmt.Println(runes[i])
	}
	fmt.Println(findFactorialWithRunes(runes))
}
