package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

func Printnumb(n int, statusChan chan string) {
	statusChan <- strconv.Itoa(rand.Intn(n))
}

func main() {
	n := 0
	fmt.Println("Chegarani tanlang")
	fmt.Scan(&n)
	statusChan := make(chan string)
	timeout := time.After(3 * time.Second)

	go Printnumb(n, statusChan)

	for {
		select {
		case numb := <-statusChan:
			fmt.Println(numb)
		case <-timeout:
			fmt.Println("time out")

		}
		if len(statusChan) == 0 {
			break
		}
	}

}
