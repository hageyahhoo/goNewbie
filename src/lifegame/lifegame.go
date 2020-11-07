package lifegame

import (
	"fmt"
)

type World struct {
	cell [][]bool
	size int
}

func createWorld(initialState [][]bool) World {
	world := new(World)

	world.size = len(initialState)
	world.cell = make([][]bool, world.size)
	for i := 0; i < world.size; i++ {
		world.cell[i] = make([]bool, world.size)
		for j := 0; j < world.size; j++ {
			world.cell[i][j] = initialState[i][j]
		}
	}

	return *world
}

func changeWorld(world World) {
	nextState := make([][]bool, world.size)
	for i := 0; i < world.size; i++ {
		world.cell[i] = make([]bool, world.size)
	}

	for i := 0; i < world.size; i++ {
		for j := 0; j < world.size; j++ {
			// TODO Refactor
			count := countNeighbors(world, i, j)
			if world.cell[i][j] {
				if count != 3 {
					nextState[i][j] = false
				}
			} else {
				if count == 2 || count == 3 {
					nextState[i][j] = true
				}
			}
		}
	}
	// FIXME
	world.cell = nextState
	fmt.Println(nextState)
}

func countNeighbors(world World, x, y int) int {
	var count int
	for i := max(x - 1, 0); i < min(world.size, x + 1); i++ {
		for j := max(y - 1, 0); j < min(world.size, y + 1); j++ {
			if world.cell[i][j] {
				count++
			}
		}
	}
	return count
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
