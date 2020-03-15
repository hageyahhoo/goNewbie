package main

import (
	"fmt"
	"runtime"
)

func main() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	fmt.Println("Alloc:", mem.Alloc)
	fmt.Println("TotalAlloc:", mem.TotalAlloc)
	fmt.Println("HeapAlloc:", mem.HeapAlloc)
	fmt.Println("NumGC:", mem.NumGC)
	fmt.Println("NumForcedGC:", mem.NumForcedGC)
}
