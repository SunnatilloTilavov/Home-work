package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

func Printnumb1(n int,c int, statusChan chan string) {
	for i:=0;i<c;i++{
	statusChan <- strconv.Itoa(rand.Intn(n))}
}

func main() {
	n := 0
	c:=0
	fmt.Println("Chegarani tanlang")
	fmt.Scan(&n)

	fmt.Println("chanel o'lchani")
	fmt.Scan(&c)

	statusChan := make(chan string,c)
	timeout := time.After(3 * time.Second)

	go Printnumb1(n,c, statusChan)

	for {
		select {
		case numb := <-statusChan:
			fmt.Println("rand numb:",numb)
		case <-timeout:
			fmt.Println("time out")

		}
		if len(statusChan) == 0 {
			break
		}
	}

}
