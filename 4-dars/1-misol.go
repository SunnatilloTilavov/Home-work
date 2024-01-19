package main

import "fmt"

func main() {
	var a,b string
	fmt.Print("1-sonni = ")
	fmt.Scan(&a)
	fmt.Print("2-sonni= ")
	fmt.Scan(&b)
	a, b = b, a
	fmt.Println("Keyingisi=", a, b)
}
