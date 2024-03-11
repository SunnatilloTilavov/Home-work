package main

import (
	"fmt"
)

func Factarial(son int) int {
	ch := make(chan int)

	go func() {
		for i := 1; i <= son; i++ {
			ch <- i
		}
		close(ch)
	}()

	s := 1
	for num := range ch {
		s *= num
	}

	return s
}

func main() {
	var son int
	fmt.Println("Son")
	fmt.Scan(&son)
	fmt.Println("Sum:", Factarial(son))
}
