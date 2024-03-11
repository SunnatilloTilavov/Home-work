////////
package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 100
	Ch := make(chan bool)
	
	go func() {
	defer close(Ch)
	for i := a; i <= b; i++ {
		Prime(i,Ch)
		if <-Ch {
			fmt.Println(a)
		}
	}
	}()

	for prime := range Ch {
		fmt.Println(prime)
	}
}

func Prime(num int, a chan<- bool)  {
	if num <= 1 {
		a<- false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			a<- false
		}
	}

	a<- true
}
