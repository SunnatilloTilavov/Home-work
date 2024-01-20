package main

import "fmt"

func main() {
	var a float32
	var b float32=3.14
	fmt.Print("radiusni kiriting = ")
	fmt.Scan(&a)
	fmt.Println("yuzi= ",a*b*b)
}