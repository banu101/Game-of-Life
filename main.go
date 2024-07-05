/*
 alem -> bootcamp -> crunch03
 contributors:
 1. nbeisenb
 2. nnurgali
 3. bgabdyzh
*/

package main

import (
	"fmt"
	"time"
)

// Variables that are used in different funcitons
var (
	liveCells   = make(map[coordinates]int)
	game        = true  // Sets false when there's errors with inputs
	done        = false // Sets true when number of live cells is zero
	stats       = false // Verbose flag
	delayMs     = 2500  // Default
	sizeInput   = true  // Input through terminal
	fullscreen  = false // FullScreen flag
	readFiles   = false // Reads from file flag
	random      = false // Random map flag
	colored     = false // Colored map flag
	edgesPortal = false // EdgesPortal map flag
	disco       = false // Flag for bonus
	emoji       = false
	filePath    = ""  // Variable that stores filepath
	height      = 0   // Variable that stores height of map
	width       = 0   // Variable that stores width of map
	footprint   = "." // Variable that stores footprint character
	dead        = "." // Variable that stores dead cell character
	live        = "×" // Variable that stores live cell character
)

// This helps for maps to have tp if not rawo keys
type coordinates struct {
	x int
	y int
}

// Main that connects everything
func main() {
	// os.Args taking flag inputs
	terminalInput()

	// If there is any error game = false
	if game {
		var inputMap [][]int
		// Reads dimension of map if not random
		if sizeInput {
			fmt.Print("Enter the height and width of the map separated by a space:\n")
			height, width = readWidthAndHeight()
			if height == -1 || width == -1 {
				fmt.Print("Error: Wrong input for Width or Height!\n")
				return
			}

		}

		// Creates map based on content from file
		if readFiles {
			inputMap = readFile(filePath)
			if inputMap == nil {
				fmt.Println("Error reading file")
				return
			}
			height = len(inputMap)
			width = len(inputMap[0])
		}

		// Reads grid values of map if not random map
		if sizeInput {
			fmt.Print("Enter ", height, " lines, each with ", width, " characters representing the grid: '.' represents an live cell, '#' represents dead cell.\n")
			inputMap = readGrid(height, width)
		}

		// Creates random map if flag random is true
		if random {
			inputMap = randomMap(height, width)
		}

		// If map is nil break the program
		if inputMap == nil {
			fmt.Print("Error: Invalid input for map\n")
			return
		}

		// Extend map if fullscreen flag is true
		if fullscreen {
			inputMap = fullScreen(inputMap, stats)
		}

		// Variable that stores amount of ticks
		tick := 1
		// Tick every 2.5 seconds
		ticker := time.Tick(time.Duration(delayMs) * time.Millisecond)

		// Goroutine to print output every 2 seconds
		go func() {
			for {
				<-ticker
				// Prints stats if verbose flag was inputted
				if stats {
					printStats(tick, height, width, delayMs)
				}
				printMap(inputMap, colored)
				// Abort mission if no cells are alive
				if done {
					game = false
				}
				calcMap(inputMap, edgesPortal)
				tick++
			}
		}()

		for game {
		}
	}
}
