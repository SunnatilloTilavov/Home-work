package main

import "fmt"

func main() {
	var a,b float32
	fmt.Print("1-tomon=")
	fmt.Scan(&a)
	fmt.Print("2-tomon=")
	fmt.Scan(&b)
	fmt.Println("Perimetri= ", 2*(a+b))
}