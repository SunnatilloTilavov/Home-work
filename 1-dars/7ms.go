package main

import "fmt"

func main() {

	var a int = 567
	c := (a % 100) % 10
	d := (a % 100) / 10

	fmt.Println("Birliklar xonasidagi son =", c)
	fmt.Println("O'nliklar xonasidagi son =", d)

}
