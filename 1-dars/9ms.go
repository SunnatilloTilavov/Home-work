package main

import "fmt"

func main() {

	var a int = 123
	c := (a % 100) % 10
	d := (a % 100) / 10
	r := a / 100
	t := c*100 + d*10 + r
	fmt.Println("Son =", t)

}
