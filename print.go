package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ANSI color codes
const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Blue  = "\033[34m"
	Green = "\033[32m"
)

// Bonus
func randomColors() string {
	num := rand.Intn(12)
	array := []string{
		"\033[90m",
		"\033[97m",
		"\033[33m",
		"\033[31m",
		"\033[30m",
		"\033[36m",
		"\033[96m",
		"\033[93m",
		"\033[95m",
		"\033[91m",
		"\033[35m",
		"\033[37m",
	}
	return array[num]
}

// Printing maps
func printMap(matrix [][]int, colored bool) {
	rand.Seed(time.Now().UnixNano())
	fmt.Print("\n")
	var cell string
	if emoji == true {
		dead = "💀"
		footprint = "👣"
		live = "😇"
	}
	for _, v := range matrix {
		for _, f := range v {
			if f == 0 {
				if disco {
					cell = randomColors() + dead + Reset
				} else if colored {
					cell = Red + dead + Reset
				} else {
					cell = dead
				}
			} else if f == 1 {
				if colored && footprint == dead {
					cell = Red + footprint + Reset
				} else if colored {
					cell = Blue + footprint + Reset
				} else {
					cell = footprint
				}
			} else if f == 2 {
				if colored {
					cell = Green + live + Reset
				} else {
					cell = live
				}
			} else {
				fmt.Print("Problem with your code in printmap!")
			}
			fmt.Print(cell + " ")
		}
		fmt.Print("\n")
	}
}

// Verbose flag
func printStats(tick, height, width, delayMs int) {
	lena := len(liveCells)
	fmt.Printf("\nTick: %d\n", tick)
	fmt.Printf("Grid Size: %dx%d\n", height, width)
	fmt.Printf("Live Cells: %d\n", lena)
	fmt.Printf("DelayMs: %d\n", delayMs)
}

// Help flag info
func helpInfo() {
	fmt.Println("Usage: go run main.go [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --help        : Show the help message and exit")
	fmt.Println("  --verbose     : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	fmt.Println("  --delay-ms=X  : Set the animation speed in milliseconds. Default is 2500 milliseconds")
	fmt.Println("  --random=HxW  : Generate a random grid of the specified width (W) and height (H)")
	fmt.Println("  --file=X      : Load the initial grid from a specified file")
	fmt.Println("  --edges-portal: Enable portal edges where cells that exit the grid appear on the opposite side")
	fmt.Println("  --fullscreen  : Adjust the grid to fit the terminal size with empty cells")
	fmt.Println("  --footprints  : Add traces of visited cells, displayed as '∘'")
	fmt.Println("  --colored     : Add color to live cells and traces if footprints are enabled")
	fmt.Println("  --disco       : Random colors for dead cells every tick")
	fmt.Println("  --emoji       : Replace default characters in emoji")
}

// Copy matrix
func copyMatrix(original [][]int) [][]int {
	rows := len(original)
	cols := len(original[0]) // Assuming all rows have the same length

	copyMatrix := make([][]int, rows)
	for i := range copyMatrix {
		copyMatrix[i] = make([]int, cols)
		copy(copyMatrix[i], original[i]) // Copy each row
	}
	return copyMatrix
}
