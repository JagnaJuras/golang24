package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Generate matrix
	matrix := generateMatrix(15, 20)

	// Print initial forest
	fmt.Println("Initial Forest:")
	// printSimpleMatrix(matrix)
	printMatrix(matrix)

	// Simulate forest fire
	treesBurned := simulateForestFire(matrix)

	// Print forest after fire
	fmt.Println("\nForest After Fire:")
	// printSimpleMatrix(matrix)
	printMatrix(matrix)

	// Calculate percentage of trees burned
	totalTrees := countTrees(matrix)
	percentBurned := float64(treesBurned) / float64(totalTrees) * 100.0

	fmt.Printf("\nPercentage of trees burned: %.2f%%\n", percentBurned)
}

func generateMatrix(rows, cols int) [][]int {
	// Create matrix
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// Fill matrix with random 0s and 1s
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Intn(2) // Generates either 0 or 1
		}
	}

	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, cell := range row {
			switch cell {
			case 0:
				fmt.Print("   ")
			case 1:
				fmt.Print("🌳 ")
			case -1:
				fmt.Print("🔥 ")
			case -2:
				fmt.Print("⚡ ")
			}
		}
		fmt.Println()
	}
}

func printSimpleMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}
}

func simulateForestFire(matrix [][]int) int {
	// Strike a random place with lightning bolt
	strikeRow := rand.Intn(len(matrix))
	strikeCol := rand.Intn(len(matrix[0]))

	matrix[strikeRow][strikeCol] = -2
	if matrix[strikeRow][strikeCol] == 0 {
		fmt.Printf("\nLightning bolt strikes at position (%d, %d), hitting an empty space\n", strikeRow, strikeCol)
	} else {
		fmt.Printf("\nLightning bolt strikes at position (%d, %d), hitting a tree\n", strikeRow, strikeCol)
	}

	// Simulate forest fire
	return igniteFire(matrix, strikeRow, strikeCol, 1)
}

func igniteFire(matrix [][]int, row, col int, probability float64) int {
	// If the cell is empty, nothing happens
	if matrix[row][col] == 0 {
		return 0
	} else if matrix[row][col] == -2 {
		matrix[row][col] = -2
	} else {
		matrix[row][col] = -1
	}

	treesBurned := 1

	// Spread fire to neighboring cells
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			// Check if the neighboring cell is within bounds
			if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) {
				// Spread fire to trees with probability
				if matrix[i][j] == 1 && rand.Float64() < probability {
					treesBurned += igniteFire(matrix, i, j, probability*0.9) // Decrease probability
				}
			}
		}
	}

	return treesBurned
}

func countTrees(matrix [][]int) int {
	count := 0
	for _, row := range matrix {
		for _, cell := range row {
			if cell == 1 {
				count++
			}
		}
	}
	return count
}
