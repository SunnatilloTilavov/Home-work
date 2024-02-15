package main

import "fmt"

func main() {
	var a float32
	var b float32=3.14
	fmt.Print("radiusni kiriting = ")
	fmt.Scan(&a)
	fmt.Printf("Yuzi= %f hajmi= %f\n",4*a*b*b,4*b/3*a*a*a)
}