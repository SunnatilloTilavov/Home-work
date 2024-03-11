///???? 16-mashq
package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	b := 50
	ch := make(chan string)
	var wg sync.WaitGroup
	go func() {
		wg.Wait()
		close(ch)
	}()

	go func(){
		for i:=10;i<=b;i++ {
		go Fibonaci(i, ch, &wg)
		if i==ch{
			fmt.Println(i)
		}
	}
	}()
}

func Fibonaci(a int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	b:=strconv.Itoa(a)
	runes := []rune(b)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	ch<-string(runes)
}

	

