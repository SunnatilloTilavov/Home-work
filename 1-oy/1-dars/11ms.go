package main

import "fmt"

func main() {

	var a int = 435
	c := (a % 100) % 10
	d := (a % 100) / 10
	r := a / 100
	t := c*100 + r*10 + d

	fmt.Println("yig'indi =", t)
}
