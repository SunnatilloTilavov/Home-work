package main

import (
	"fmt"
)

func Print123() int {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	s := 0
	for num := range ch {
		s += num
	}

	return s
}

func main() {
	fmt.Println("Sum:", Print123())
}
