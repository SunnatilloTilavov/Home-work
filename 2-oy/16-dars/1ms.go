package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch:=make(chan int,100)

	go func(){
		for i := 1; i <= 100; i++ {
		c:=rand.Intn(1000)
		ch<- c
	}
}()
for val:=range ch{
	fmt.Println(val)
	if len(ch)==0{
		break
	}
}
}

