package main

import (
	"fmt"
)

func Reaverse(soz string, ch chan string) {
	go func() {

		runes := []rune(soz)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		ch<-string(runes)
	}()
	
}

func main() {
	ch := make(chan string)
	var soz string
	fmt.Println("string:")
	fmt.Scan(&soz)
	Reaverse(soz, ch)
	fmt.Println(<-ch)
}
