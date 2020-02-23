package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("You need to pass 2 arguments. HP and Attack.")
		os.Exit(9)
	}

	var (
		hp     int
		attack int
		result int
	)

	/*
	 * flag.Arg(): Start from 0
	 * os.Args: Start from 1
	 */
	hp, _ = strconv.Atoi(flag.Arg(0))
	attack, _ = strconv.Atoi(flag.Arg(1))
	result = hp / attack
	if hp%attack != 0 {
		result++
	}

	fmt.Println("HP:", hp, "Attack:", attack, "Result:", result)
}
