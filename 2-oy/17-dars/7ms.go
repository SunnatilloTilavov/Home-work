package main

import (
	"fmt"
)

func Maxnumb(s []int,ch chan int) {

	var max int
	go func() {

		max=s[0]
		for i:=0;i<len(s);i++{
		if max<s[i]{
			max=s[i]
		}	
		ch<-max
	}	
	}()

}

func main() {
	ch := make(chan int)
	var son,n int
	fmt.Println("slise count:")
	fmt.Scan(&n)
	s:= make([]int, n)
	for i:=0;i<n;i++{
		fmt.Print("Son:")
		fmt.Scan(&son)
		s[i]=son
	}
	Maxnumb(s,ch)
	n=<-ch
	fmt.Println("MAX:",n )
}