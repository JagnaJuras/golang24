package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// ----------------- ONE SIM ------------------
	// Generate matrix
	matrix := generateMatrix(15, 20, 0.5)

	fmt.Println("Initial Forest:")
	printMatrix(matrix)

	generateImage(matrix, "before.jpg")

	// Simulate forest fire
	simulateForestFire(matrix)

	fmt.Println("Forest After Fire:")
	printMatrix(matrix)

	generateImage(matrix, "after.jpg")

	// Calculate percentage of trees burned
	percentBurned := countBurntTrees(matrix)
	fmt.Printf("\nPercentage of trees burned: %.2f%%\n", percentBurned)
	// ------------ TESTING DIFFERENT DENSITIES WITH ALL FIRES -------------

	rows := []int{10, 50, 100, 500, 1000, 50, 100, 200}
	columns := []int{10, 10, 10, 10, 10, 50, 100, 200}
	densities := []float64{0.4, 0.45, 0.5, 0.55, 0.6, 0.65, 0.7, 0.75, 0.8, 0.85, 0.9}
	numSimulations := 1000

	pAllFires := plot.New()

	pAllFires.Title.Text = "Average Percentage of Trees Burned (All Fires)"
	pAllFires.X.Label.Text = "Tree Density"
	pAllFires.Y.Label.Text = "Average Percentage Burned"

	var linesAllFires []*plotter.Line

	for i := range rows {
		results := make([]float64, len(densities))
		for j, density := range densities {
			results[j] = runSimulations(rows[i], columns[i], density, numSimulations)
		}

		points := make(plotter.XYs, len(densities))
		for j := range densities {
			points[j].X = densities[j]
			points[j].Y = results[j]
		}

		line, err := plotter.NewLine(points)
		if err != nil {
			fmt.Printf("Error creating line plot: %v\n", err)
			return
		}
		line.LineStyle.Width = vg.Points(2)
		line.LineStyle.Dashes = []vg.Length{}
		line.Color = getRandomColor()
		linesAllFires = append(linesAllFires, line)
	}

	for i, line := range linesAllFires {
		pAllFires.Add(line)
		pAllFires.Legend.Add(fmt.Sprintf("%d x %d", rows[i], columns[i]), line)
	}

	// Set legend position
	pAllFires.Legend.Top = true
	pAllFires.Legend.Left = true
	pAllFires.Legend.YOffs = -10

	// Generate plot file for all fires
	plotFilenameAllFires := "forest_fire_plot_all_fires.png"

	if err := pAllFires.Save(8*vg.Inch, 6*vg.Inch, plotFilenameAllFires); err != nil {
		fmt.Printf("Error saving plot (All Fires): %v\n", err)
	}

	// ------------ TESTING DIFFERENT DENSITIES WITH FIRES WHERE TREES WERE HIT -------------

	pTreeFires := plot.New()

	pTreeFires.Title.Text = "Average Percentage of Trees Burned (Fires Where Trees Were Hit)"
	pTreeFires.X.Label.Text = "Tree Density"
	pTreeFires.Y.Label.Text = "Average Percentage Burned"

	var linesTreeFires []*plotter.Line

	for i := range rows {
		results := make([]float64, len(densities))
		for j, density := range densities {
			results[j] = runSimulationsWithFires(rows[i], columns[i], density, numSimulations)
		}

		points := make(plotter.XYs, len(densities))
		for j := range densities {
			points[j].X = densities[j]
			points[j].Y = results[j]
		}

		line, err := plotter.NewLine(points)
		if err != nil {
			fmt.Printf("Error creating line plot: %v\n", err)
			return
		}
		line.LineStyle.Width = vg.Points(2)
		line.LineStyle.Dashes = []vg.Length{}
		line.Color = getRandomColor()
		linesTreeFires = append(linesTreeFires, line)
	}

	for i, line := range linesTreeFires {
		pTreeFires.Add(line)
		pTreeFires.Legend.Add(fmt.Sprintf("%d x %d", rows[i], columns[i]), line)
	}

	// Set legend position
	pTreeFires.Legend.Top = true
	pTreeFires.Legend.Left = true
	pTreeFires.Legend.YOffs = -10

	// Generate plot file for fires where trees were hit
	plotFilenameTreeFires := "forest_fire_plot_tree_fires.png"

	if err := pTreeFires.Save(8*vg.Inch, 6*vg.Inch, plotFilenameTreeFires); err != nil {
		fmt.Printf("Error saving plot (Tree Fires): %v\n", err)
	}
}

func getRandomColor() color.RGBA {
	return color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255}
}

func runSimulations(rows, columns int, treeDensity float64, numSimulations int) float64 {
	totalPercentBurned := 0.0

	for i := 0; i < numSimulations; i++ {
		matrix := generateMatrix(rows, columns, treeDensity)
		simulateForestFire(matrix)
		percentBurned := countBurntTrees(matrix)
		totalPercentBurned += percentBurned
	}

	return totalPercentBurned / float64(numSimulations)
}

func runSimulationsWithFires(rows, columns int, treeDensity float64, numSimulations int) float64 {
	totalPercentBurned := 0.0
	totalFires := 0

	for i := 0; i < numSimulations; i++ {
		matrix := generateMatrix(rows, columns, treeDensity)
		simulateForestFire(matrix)
		percentBurned := countBurntTrees(matrix)
		if percentBurned > 0 {
			totalPercentBurned += percentBurned
			totalFires++
		}
	}

	// Calculate average percentage burned only for cases where a fire started
	averagePercentBurned := 0.0
	if totalFires > 0 {
		averagePercentBurned = totalPercentBurned / float64(totalFires)
	}

	return averagePercentBurned
}

func generateMatrix(rows, cols int, treePercentage float64) [][]int {
	// Create matrix
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// Calculate total number of cells that should contain trees
	totalCells := rows * cols
	totalTrees := int(float64(totalCells) * treePercentage)

	// Randomly select cells to contain trees
	treeCells := make(map[int]bool)
	for totalTrees > 0 {
		cell := rand.Intn(totalCells)
		if !treeCells[cell] {
			treeCells[cell] = true
			totalTrees--
		}
	}

	// Fill matrix based on selected tree cells
	for cell := range treeCells {
		row := cell / cols
		col := cell % cols
		matrix[row][col] = 1 // Place a tree
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

func simulateForestFire(matrix [][]int) {
	// Strike a random place with lightning bolt
	strikeRow := rand.Intn(len(matrix))
	strikeCol := rand.Intn(len(matrix[0]))

	if matrix[strikeRow][strikeCol] == 0 {
		// fmt.Printf("\nLightning bolt strikes at position (%d, %d), hitting an empty space\n", strikeRow, strikeCol)
		matrix[strikeRow][strikeCol] = -2
		return
	} else {
		// fmt.Printf("\nLightning bolt strikes at position (%d, %d), hitting a tree\n", strikeRow, strikeCol)
		matrix[strikeRow][strikeCol] = -2
		igniteFire(matrix, strikeRow, strikeCol, 1)
	}
}

func igniteFire(matrix [][]int, row, col int, probability float64) {
	if matrix[row][col] == -2 {
		matrix[row][col] = -2
	} else {
		matrix[row][col] = -1
	}

	// Spread fire to neighboring cells
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			// Check if the neighboring cell is within bounds
			if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) {
				// Spread fire to trees with probability
				if matrix[i][j] == 1 && rand.Float64() < probability {
					igniteFire(matrix, i, j, probability*1) // Decrease probability
				}
			}
		}
	}

}

func countBurntTrees(matrix [][]int) float64 {
	totalTrees := 0
	burntTrees := 0

	for _, row := range matrix {
		for _, cell := range row {
			if cell == 1 {
				totalTrees++
			} else if cell == -1 {
				burntTrees++
			}
		}
	}

	return float64(burntTrees) / (float64(totalTrees) + float64(burntTrees)) * 100
}

func generateImage(matrix [][]int, filename string) error {
	emptyColor := color.RGBA{255, 255, 255, 255} // White
	treeColor := color.RGBA{34, 139, 34, 255}    // Forest green
	fireColor := color.RGBA{255, 0, 0, 255}      // Red
	lightningColor := color.RGBA{0, 0, 0, 255}   // Yellow

	rows := len(matrix)
	cols := len(matrix[0])
	imageSize := image.Rect(0, 0, cols, rows)

	img := image.NewRGBA(imageSize)

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			switch matrix[y][x] {
			case 0:
				img.Set(x, y, emptyColor)
			case 1:
				img.Set(x, y, treeColor)
			case -1:
				img.Set(x, y, fireColor)
			case -2:
				img.Set(x, y, lightningColor)
			}
		}
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		return err
	}

	return nil
}
