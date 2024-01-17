package main

import "fmt"

func main() {

	var a int = 234
	c := (a % 100) % 10
	d := (a % 100) / 10
	r := a / 100

	fmt.Println("yig'indi =", c+d+r)

}
