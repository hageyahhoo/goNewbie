package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()
	var (
		hp     int
		attack int
		result int
	)

	// flag.Arg(): Start from 0
	// os.Args: Start from 1
	hp, _ = strconv.Atoi(flag.Arg(0))
	attack, _ = strconv.Atoi(flag.Arg(1))
	result = hp / attack
	if hp%attack != 0 {
		result++
	}

	fmt.Println("HP:", hp, "Attack:", attack, "Result:", result)
}
