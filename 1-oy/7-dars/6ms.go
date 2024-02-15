package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("A = ")
	fmt.Scan(&a)

	fmt.Print("B = ")
	fmt.Scan(&b)

	for a != b {
		if a > b {
			a %= b
		} else {
			b %= a
		}

		if a == 0 {
			a = b
		}
		if b == 0 {
			b = a
		}
	}

	fmt.Println(a)
}
