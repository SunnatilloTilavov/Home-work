package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func Printnumb(n int, statusChan chan string) {
	statusChan <- strconv.Itoa(rand.Intn(n))
}

func main() {
	n := 0
	fmt.Println("Chegarani tanlang")
	fmt.Scan(&n)
	statusChan := make(chan string)

	go Printnumb(n, statusChan)

	numb := <-statusChan
	fmt.Println(numb)

}
