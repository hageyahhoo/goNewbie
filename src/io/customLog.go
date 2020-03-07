package main

import (
	"fmt"
	"log"
	"os"
)

var LOG_FILE = "./custom.log"

func main() {
	f, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(9)
	}
	defer f.Close()

	customLog := log.New(f, "CustomLog", log.LstdFlags|log.Lshortfile)
	//	customLog.Fatal("OH!")
	//	customLog.Panic("orz")
	customLog.Println("Los Ingobernables de Japon!")
}
