package main

import "fmt"

func main() {

	a := 5
	b := 7
	a, b = b, a

	fmt.Println("Keyingisi=", a, b)
}
