package main

import "fmt"

func main() {
	var a,b,c,d float32
	fmt.Print("1-tomon=")
	fmt.Scan(&a)
	fmt.Print("2-tomon=")
	fmt.Scan(&b)
	fmt.Print("3-tomon=")
	fmt.Scan(&c)
	fmt.Print("4-tomon= ")
	fmt.Scan(&d)
	fmt.Println("Perimetri= ", a+b+c+d)
}