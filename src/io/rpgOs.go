package main

import (
	"fmt"
	"os"
	"strconv"
)

// TODO Add UT
func main() {
	var (
		hp     int
		attack int
		result int
	)

	// TODO Add error handling & logging
	// TODO Check convert error
	// TODO Check length error of Args
	hp, _ = strconv.Atoi(os.Args[1])
	attack, _ = strconv.Atoi(os.Args[2])
	result = hp / attack
	if hp%attack != 0 {
		result++
	}

	fmt.Println("HP:", hp, "Attack:", attack, "Result:", result)
}
