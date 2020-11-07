package lifegame

import (
	"testing"
)

func TestCreateWorld(t *testing.T) {
	// TODO Change param from indices to state
	world := createWorld(3, 3)
	expected := [3][3]bool {
		{ false, false, false },
		{ false, false, false },
		{ false, false, false },
	}

	for i, cells := range expected {
		for j, cell := range cells {
			if world.cell[i][j] != cell {
				t.Errorf("Cell error: [%d][%d]", i, j)
			}
		}
	}
}
