package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup


	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go calculate(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	
	sum := 0
	for square := range ch {
		sum += square
	}
	fmt.Println("Sum", sum)
}

func calculate(number int, a chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	b := number * number
	a <- b
}