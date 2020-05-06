package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int64 = 5
	var p1 = &a
	var p2 = (*int32)(unsafe.Pointer(p1))
	// We can access values of each pointer by using "*"
	fmt.Println(*p1)
	fmt.Println(*p2)

	// Lead overflow
	*p1 = 54432342334355353
	fmt.Println(*p1)
	fmt.Println(*p2)
}
