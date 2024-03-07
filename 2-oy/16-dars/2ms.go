package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomch, readablech := sendrandm()
	for {
		select {
		case num := <-randomch:
			fmt.Println(num)
			// time.Sleep(1 *time.Second)
		case status := <-readablech:
			fmt.Print(status)
		}
	}
}

func sendrandm() (chan int, <-chan string) {
	randomch := make(chan int)
	readablech := make(chan string)
	go func() {
		defer close(randomch)
		defer close(readablech)
		for {
			select {
			case randomch <- rand.Intn(100):
				readablech <- "NUM:"
			case <-time.After(time.Second):
				readablech <- "time out"
				return
			}

		}
	}()
	return randomch, readablech
}
