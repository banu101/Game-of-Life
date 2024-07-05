package main

import (
	"fmt"
	"os"
	"strings"
)

// Only for reading height and width
func readWidthAndHeight() (int, int) {
	var width, height int

	_, err := fmt.Scanf("%d %d", &height, &width)
	if err != nil || (height < 3 || width < 3) {
		return -1, -1
	}

	return height, width
}

// Scans h bumber of lines each with w number of characters
func readGrid(height, width int) [][]int {
	var cell rune
	inputMap := make([][]int, height)
	liveCells = make(map[coordinates]int)

	for i := 0; i < height; i++ {
		inputMap[i] = make([]int, width)
		for j := 0; j < width; j++ {
			if _, err := fmt.Scanf("%c", &cell); err != nil {
				return nil
			}

			if cell != '.' && cell != '#' {
				return nil
			}

			if cell == '#' {
				liveCells[coordinates{x: i, y: j}] = 1 // Here, 1 only means that key is not empty
				inputMap[i][j] = 2
			}

		}

		if _, err := fmt.Scanf("%c", &cell); err != nil {
			return nil
		}
	}

	if len(liveCells) == 0 {
		return nil
	}

	return inputMap
}

// Reads a map from a specified file
func readFile(filePath string) [][]int {
	// Attempts to read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	// Convert the file content to a string and remove any carriage return characters "\r"
	lines := strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n")

	for i := 0; i < len(lines)-1; i++ {
		if len(lines[i]) != len(lines[i+1]) {
			fmt.Println("Error: All lines must have the same length")
			return nil
		}
	}

	height = len(lines)
	width = len(lines[0])

	inputMap := make([][]int, height)
	for i := range inputMap {
		inputMap[i] = make([]int, width)
	}

	liveCells = make(map[coordinates]int)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cell := rune(lines[i][j])
			if cell != '.' && cell != '#' {
				fmt.Println("Error: Invalid character in the file")
				return nil
			}

			if cell == '#' {
				liveCells[coordinates{x: i, y: j}] = 1 // Here, 1 only means that key is not empty
				inputMap[i][j] = 2
			}
		}
	}

	return inputMap
}
