package main

import (
	"fmt"
	"strconv"
	"time"
)

func Printnumb3( statusChan chan string) {
	for i:=1;i<=10;i++{
	statusChan <-strconv.Itoa(i)}
}

func main() {
	statusChan := make(chan string)
	timeout := time.After(1 * time.Second)

	go Printnumb3( statusChan)
	stop:=false

	for {
		select {
		case numb := <-statusChan:
			fmt.Println(numb)
		case <-timeout:
		stop=true

		}
		if stop {
			break
		}
	}

}