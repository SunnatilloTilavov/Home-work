package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	ch := make(chan float64) 
	var wg sync.WaitGroup
	numbers := [5]int{1, 2, 3, 4, 5}

	for _, num := range numbers { 
		wg.Add(1)
		go Sqrt(num, ch, &wg) 
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for sqrt := range ch {
		fmt.Println("Sqrt", sqrt)
	}
}

func Sqrt(number int, ch chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result := math.Sqrt(float64(number))
	ch <- result
}
