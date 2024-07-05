package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tsize "github.com/kopoli/go-terminal-size"
)

// Iterates over the entire game of life matrix to calculate the next generation of alive cells
func calcMap(matrix [][]int, edgesPortal bool) {
	matrixOld := copyMatrix(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			countAdjacentCells(matrix, matrixOld, i, j, edgesPortal)
		}
	}
}

// Counts number of live cells in adjacent cells
func countAdjacentCells(matrix, matrixOld [][]int, x, y int, edgesPortal bool) {
	directions := [8][2]int{
		{-1, -1}, // Top left
		{-1, 0},  // Top
		{-1, 1},  // Top right
		{0, -1},  // Left
		{0, 1},   // Right
		{1, -1},  // Bottom left
		{1, 0},   // Bottom
		{1, 1},   // Bottom right
	}
	count := 0
	for _, dir := range directions {
		// Calculate the new coordinates based on the current cell and direction
		newX, newY := x+dir[0], y+dir[1]

		// Checks if edgesPortak flag is true
		if edgesPortal {
			// If the coordinate on the x-axis is less than zero, then moving to the extreme point
			if newX < 0 {
				newX = len(matrixOld) - 1
				// If the coordinate on the x-axis is greater than length, then moving to zero point
			} else if newX >= len(matrixOld) {
				newX = 0
			}
			// If the coordinate on the y-axis is less than zero, then moving to the extreme point
			if newY < 0 {
				newY = len(matrixOld[0]) - 1
				// If the coordinate on the y-axis is greater than length, then moving to zero point
			} else if newY >= len(matrixOld[0]) {
				newY = 0
			}
			// Checks the cell at that position is alive and increment the counter
			if matrixOld[newX][newY] == 2 {
				count++
			}
		} else {
			// Checks if the new coordinates are within the bounds and the cell at that position is alive and increment the counter
			if newX >= 0 && newX < len(matrixOld) && newY >= 0 && newY < len(matrixOld[0]) && matrixOld[newX][newY] == 2 {
				count++
			}
		}
	}
	// Any live cell with fewer than two live neighbors dies (underpopulation) or more than three live neighbors dies (overpopulation)
	if count != 2 && count != 3 && matrixOld[x][y] == 2 {
		matrix[x][y] = 1
		delete(liveCells, coordinates{x: x, y: y})
		// Any live cell with two or three live neighbors lives on to the next generation and any dead cell with exactly three live neighbors becomes a live cell (reproduction).
	} else if count == 3 && (matrixOld[x][y] == 0 || matrixOld[x][y] == 1) {
		matrix[x][y] = 2
		liveCells[coordinates{x: x, y: y}] = 1
	}

	// If there are no living cells left on the grid, the game ends
	if len(liveCells) == 0 {
		done = true
	}
}

// Inputtting flags
func terminalInput() {
	input := os.Args[1:]
	var flag []rune
	for _, v := range input {
		flag = []rune(v)
		// Last to conditions are to make sure there is no conflict with previous flags
		if len(flag) == 6 && v == "--help" {
			if !stats && delayMs == 2500 && !random && !fullscreen && !colored && !edgesPortal && !readFiles && footprint != "∘" {
				helpInfo()
				game = false
				return
			}
			// Verbose flag
		} else if len(flag) == 9 && v == "--verbose" {
			stats = true
			// Delay-ms flag
		} else if len(flag) > 11 && string(flag[:11]) == "--delay-ms=" {
			number, err := strconv.Atoi(string(flag[11:]))
			if err != nil || number <= 0 {
				game = false
				fmt.Println("Invalid delay input!")
				return
			}
			delayMs = number
			// Random flag
		} else if len(flag) > 9 && string(flag[:9]) == "--random=" {
			if !readFiles {
				numbers := strings.Split(string(flag[9:]), "x")
				if len(numbers) == 2 {
					number0, err0 := strconv.Atoi(numbers[0])
					if err0 != nil {
						game = false
						fmt.Println("Invalid input size for a random map!")
						return
					}
					height = number0

					number1, err1 := strconv.Atoi(numbers[1])
					if err1 != nil {
						game = false
						fmt.Println("Invalid input size for a random map!")
						return
					}
					width = number1
				} else {
					game = false
					fmt.Println("Invalid input size for a random map!")
					return
				}
				random = true
				sizeInput = false
			}
			// Fullscreen flag
		} else if len(flag) == 12 && v == "--fullscreen" {
			_, err := tsize.GetSize()
			if err == nil {
				fullscreen = true
			}
			if emoji {
				fullscreen = false
			}
			// Footprints flag
		} else if len(flag) == 12 && v == "--footprints" {
			footprint = "∘"
			// Colored flag
		} else if len(flag) == 9 && v == "--colored" {
			colored = true
			// Read from file flag
		} else if len(flag) > 7 && string(flag[:7]) == "--file=" {
			if !random {
				_, err := os.ReadFile(string(flag[7:]))
				if err != nil {
					fmt.Println("crunch03: " + string(flag[7:]) + ": No such directory or file")
				}
				readFiles = true
				sizeInput = false
				filePath = string(flag[7:])
			}
			// Edges portal flag
		} else if len(flag) == 14 && v == "--edges-portal" {
			edgesPortal = true
		} else if len(flag) == 7 && v == "--disco" {
			disco = true
		} else if len(flag) == 7 && v == "--emoji" {
			emoji = true
			if fullscreen {
				emoji = false
			}
		} else {
			game = false
			fmt.Println("Invalid flag!")
			return
		}
	}
}
