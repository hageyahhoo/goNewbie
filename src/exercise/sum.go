package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var LOG_FILE = "./custom.log"

func main() {
	// FIXME How to handle defer + Close by calling other small functions?
	file, err1 := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Println("Failed to open a log file!")
		os.Exit(9)
	}
	defer file.Close()

	customLog := log.New(file, "SUM", log.LstdFlags|log.Lshortfile)

	// FIXME variable name of error(s)
	summary, err2 := getSummary(os.Args)
	if err2 != nil {
		customLog.Println(err2.Error())
	}

	fmt.Println("Summary:", summary)
}

func getSummary(args []string) (int, error) {
	length := len(args)
	if length == 1 {
		return 0, errors.New("Please pass numbers!")
	}

	var summary int
	for i := 1; i < length; i++ {
		v, e := strconv.Atoi(args[i])
		if e != nil {
			return 0, errors.New("Passed value is not number!")
		}
		summary += v
	}

	return summary, nil
}
