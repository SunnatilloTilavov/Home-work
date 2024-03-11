package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int) 
	var wg sync.WaitGroup
	numbers := [7]int{1, 2, 3, 4, 5,6,8}

	for _, num := range numbers { 
		wg.Add(1)
		go sum(num, ch, &wg) 
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	sum1:=0
	for s := range ch {
		sum1+=s
	}
	fmt.Println(sum1)
}

func sum(number int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	if number%2==0{
		ch<-number
	}else{
		ch<-0
	}
	
}
