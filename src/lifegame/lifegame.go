package lifegame

type World struct {
	cell [][]bool
}

func createWorld(x, y int) World {
	w := new(World)

	w.cell = make([][]bool, x)
	for i := 0; i < x; i++ {
		w.cell[i] = make([]bool, y)
	}

	return *w
}
