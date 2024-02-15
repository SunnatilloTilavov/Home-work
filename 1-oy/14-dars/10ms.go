package main

import (
	"fmt"
)

func main() {
	var n, son int
	fmt.Print("massiv nechta sondan iborat=")
	fmt.Scan(&n)
	slice := make([]int, n)
	fmt.Println(slice)
	for i := 0; i < n; i++ {
		fmt.Print("son=")
		fmt.Scan(&son)
		slice[i] = son
	}
	a := 0
	b := 0
	c := 0
	for i := 0; i < n; i++ {
		if slice[i] > 0 {
			a += 1
		} else if slice[i] == 0 {
			b += 1
		} else {
			c += 1
		}
	}
	if a > 0 {
		fmt.Println("Possitive=", float32(a)/float32(n))
	}
	if c > 0 {
		fmt.Println("Negativ=", float32(c)/float32(n))
	}
	if b > 0 {
		fmt.Println("Zero=", float32(b)/float32(n))
	}
}
