package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	Text:="asdsd sd asd sadasd ad  adasda"
	Ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go count(Text, Ch, &wg)
	go func() {
		wg.Wait()
		close(Ch)
	}()
	Count := <-Ch
	fmt.Printf("Count: %d\n", Count)
}

func count(text string, Ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()	
	Ch <- len( strings.Fields(text))
}
