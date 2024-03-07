package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		go NumberPrint(i, &wg)
		wg.Add(1)
	}
	wg.Wait()
}

func NumberPrint(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(a)
}
