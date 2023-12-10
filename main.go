// Define a type to represent a slice of slices that each hold a boolean value
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type World [][]bool

// Method to print the world
func (w World) Print() {
	for _, row := range w {
		for _, cell := range row {
			if cell {
				fmt.Print("🟩 ")
			} else {
				fmt.Print("🟫 ")
			}
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

	fmt.Print("Enter the width of the world: ")
	fmt.Scan(&width)

	fmt.Print("Enter the height of the world: ")
	fmt.Scan(&height)

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
		world.Print()
		time.Sleep(100 * time.Millisecond)
		world.Step()
		fmt.Print("\033[H\033[2J")
	}
}
