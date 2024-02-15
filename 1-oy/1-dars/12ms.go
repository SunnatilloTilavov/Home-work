package main

import "fmt"

func main() {

	var a int = 123
	c := (a % 100) % 10
	d := (a % 100) / 10
	r := a / 100
	t := d*100 + r*10 + c

	fmt.Println("yig'indi =", t)

}
