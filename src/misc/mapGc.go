package main

import (
	"runtime"
)

func main() {
	var N = 10000
	myMap := make(map[int]*int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = &value
	}
	runtime.GC()
	_ = myMap[0]
}
