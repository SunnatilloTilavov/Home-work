package main

import "fmt"

func main() {

	a, b, c := 1, 2, 3
	fmt.Println("Natija=", a, b, c)

	a, b, c = b, c, a

	fmt.Println("Natija=", a, b, c)
}
