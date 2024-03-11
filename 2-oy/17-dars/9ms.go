package main

import (
	"fmt"
)

func multiplyBy2(n int, son []int) []int {
	ch := make(chan int)
	go func() {
		for _, i := range son {
			ch <- i* 2
		}
		close(ch)
	}()

	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = <-ch
	}

	return s
}

func main() {
	var son, n int

	fmt.Print("Slice count: ")
	fmt.Scan(&n)

	s := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Son:")
		fmt.Scan(&son)
		s[i] = son
	}

	result := multiplyBy2(n, s)
	fmt.Println(result)
}
