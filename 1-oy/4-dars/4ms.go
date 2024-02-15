package main

import "fmt"

func main() {
	var a float32
	p:=3.14
	fmt.Print("Diametri=")
	fmt.Scan(&a)
	fmt.Println("Perimetri= %f",p*a)
}