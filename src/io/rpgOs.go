package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// TODO Add UT
func main() {
	// TODO Add error handling & logging
	if len(os.Args) != 3 {
		fmt.Println("You need to pass 2 arguments. HP and Attack.")
		os.Exit(9)
	}

	var (
		hp     int
		attack int
		result int
	)

	// TODO Check convert error
	hp, _ = strconv.Atoi(os.Args[1])
	attack, _ = strconv.Atoi(os.Args[2])
	result = hp / attack
	if hp%attack != 0 {
		result++
	}

	fmt.Println("HP:", hp, "Attack:", attack, "Result:", result)
	io.WriteString(os.Stdout, "DONE\n")
	io.WriteString(os.Stderr, "Fin.\n")
}
