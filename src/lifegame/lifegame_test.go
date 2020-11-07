package lifegame

import (
	"testing"
)

// TODO print as string
func TestCreateWorld(t *testing.T) {
	initialState := [][]bool {
		{ false, true, false },
		{ false, true, false },
		{ false, true, false },
	}
	world := createWorld(initialState)

	for i, cells := range initialState {
		for j, cell := range cells {
			if world.cell[i][j] != cell {
				t.Errorf("Cell error: [%d][%d] != %t", i, j, cell)
			}
		}
	}
}


func TestChangeWorld(t *testing.T) {
	initialState := [][]bool {
		{ false, true, false },
		{ false, true, false },
		{ false, true, false },
	}
	world := createWorld(initialState)

	changeWorld(world)

	nextState := [][]bool {
		{ false, false, false },
		{ true, true, true },
		{ false, false, false },
	}
	for i, cells := range nextState {
		for j, cell := range cells {
			if world.cell[i][j] != cell {
				t.Errorf("Cell error: [%d][%d] != %t", i, j, cell)
			}
		}
	}
}
