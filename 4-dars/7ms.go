package main

import "fmt"

func main() {
	var a float32
	fmt.Print("tomoni=")
	fmt.Scan(&a)
	fmt.Println("Perimetri= %f", a*4)
}