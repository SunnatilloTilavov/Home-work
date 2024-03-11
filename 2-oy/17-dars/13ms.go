package main

import (
	"fmt"
	"sync"
)

func main() {	
	a := 1
	b := 1000
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := a; i <= b; i++ {
		wg.Add(1)
		go calculate1(i, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	sum := 0
	for numb := range ch {
		sum += numb
	}
	fmt.Println("Sum ", a, "to",b, ":", sum)
}

func calculate1(number int, a chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	b := number * number
	a <- b
}
