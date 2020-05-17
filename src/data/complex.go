package main

import (
	"fmt"
)

func main() {
	c1 := 12 + 5i
	c2 := complex(5, 7)

	fmt.Printf("Type of c1 : %T\n", c1)
	fmt.Printf("Value of c1 : %v\n", c1)
	fmt.Printf("Type of c2 : %T\n", c2)
	fmt.Printf("Value of c2 : %v\n", c2)

	fmt.Println("c1 + c2 =", c1+c2)
	fmt.Println("c1 - c2 =", c1-c2)
	fmt.Println("c1 * c2 =", c1*c2)
	fmt.Println("c1 / c2 =", c1/c2)
}
