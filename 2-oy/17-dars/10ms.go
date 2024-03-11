package main

import (
	"fmt"
)

func AVG(son []int) int {
	ch := make(chan int)
	var s int 

	go func() {
		for i := 0; i < len(son); i++ {
			ch<-son[i]
		}
		close(ch)
	}()

	for num := range ch {
		s += num
	}

	return s/len(son)
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

	fmt.Println("AVG:", AVG(s))
}
