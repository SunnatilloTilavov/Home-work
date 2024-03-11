package main

import (
	"fmt"
)

func Primenumber(son int) bool{
	ch := make(chan bool)

	go func() {
		defer close(ch)
		if son==1||son==2||son==0{
		ch<-false
		}else if son%2==0{
			ch<-false
		}
		for i := 3; i < son; i += 2 {
			if son % i == 0 {
				ch<-false
			}
		}
		ch<-true	
	}()

	return <-ch
}

func main() {
	var son int
	fmt.Print("Son:")
	fmt.Scan(&son)
	fmt.Println("Tub son:", Primenumber(son))
}
