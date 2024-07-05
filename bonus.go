package main

import (
	"fmt"
	"math/rand"

	tsize "github.com/kopoli/go-terminal-size"
)

// Adjust the map to fit the terminal size in fullscreen mode
func fullScreen(matrix [][]int, stats bool) [][]int {
	// Get the terminal size
	size, err := tsize.GetSize()

	var (
		termHeight int
		termWidth  int
	)

	if err != nil {
		fmt.Println("Error occured while receiving the terminal")
		return matrix
	}

	termHeight = size.Height - 1
	termWidth = (size.Width) / 2

	// Adjust height if --verbose mode is enabled
	if stats {
		termHeight -= 5
	}

	newMatrix := make([][]int, termHeight)
	for i := range newMatrix {
		newMatrix[i] = make([]int, termWidth)
	}

	// Copy the original matrix data to the new matrix within the usable dimensions
	for i := 0; i < len(matrix) && i < termHeight; i++ {
		for j := 0; j < len(matrix[i]) && j < termWidth; j++ {
			newMatrix[i][j] = matrix[i][j]
		}
	}

	return newMatrix
}

// Generation of a random map in Case 2 above
func randomMap(h, w int) [][]int {
	if h < 3 || w < 3 {
		return nil
	}

	m := make([][]int, h)

	for i := 0; i < h; i++ {
		m[i] = make([]int, w)
	}

	var x, y int

	for len(liveCells) < (h * w / 3) {
		x, y = generateRandomCoordinates(h, w)
		m[x][y] = 2
	}

	return m
}

// Putting live cells in a random cell
func generateRandomCoordinates(height, width int) (int, int) {
	next := coordinates{rand.Intn(height - 1), rand.Intn(width - 1)}
	if liveCells[coordinates{next.x, next.y}] == 0 {
		liveCells[coordinates{next.x, next.y}] = 1
	}

	return next.x, next.y
}
