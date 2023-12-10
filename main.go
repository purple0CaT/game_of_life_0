// Define a type to represent a slice of slices that each hold a boolean value
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type World [][]bool

// Method to print the world
func getTheme(theme string, val bool) int {
	res := "⬜️"

	if theme == "night" {
		if !val {
			res = "⬛️"
		}
	} else {
		if val {
			res = "⬛️"
		}
	}
	fmt.Print(res)
	return 0
}

func (w World) Print(theme string) {
	for _, row := range w {
		for _, cell := range row {
			getTheme(theme, !!cell)
		}
		fmt.Println()
	}
}

// Method to simulate the game
func (w World) Step() {
	// Create a new world with the same dimensions as the current world
	newWorld := make(World, len(w))
	for i := range newWorld {
		newWorld[i] = make([]bool, len(w[i]))
	}

	// Iterate over each cell in the world
	for i, row := range w {
		for j := range row {
			// Count the number of live neighbors
			liveNeighbors := 0
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					if i+x < 0 || i+x >= len(w) {
						continue
					}
					if j+y < 0 || j+y >= len(row) {
						continue
					}
					if w[i+x][j+y] {
						liveNeighbors++
					}
				}
			}

			// Apply the rules of the game
			if w[i][j] {
				if liveNeighbors == 2 || liveNeighbors == 3 {
					newWorld[i][j] = true
				}
			} else {
				if liveNeighbors == 3 {
					newWorld[i][j] = true
				}
			}
		}
	}

	// Update the current world
	for i := range w {
		copy(w[i], newWorld[i])
	}
}

func main() {
	var width, height int
	var theme string

	fmt.Print("Enter the width of the world: ")
	fmt.Scan(&width)

	fmt.Print("Enter the height of the world: ")
	fmt.Scan(&height)

	fmt.Print("Enter the theme. (day/night): ")
	fmt.Scan(&theme)

	// Create a new world with random initial values
	world := make(World, height)
	for i := range world {
		world[i] = make([]bool, width)
		for j := range world[i] {
			world[i][j] = rand.Intn(2) == 1
		}
	}

	// Simulate the game
	for {
		world.Print(theme)
		time.Sleep(100 * time.Millisecond)
		world.Step()
		fmt.Print("\033[H\033[2J")
	}
}
