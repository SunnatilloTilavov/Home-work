package main

import (
	"fmt"
)

func Sort(son []int) []int {
	// ch := make(chan int) 
	go func() {
		for i:=0;i<len(son);i++{
			for j:=0;j<len(son);j++{
				if son[i]<son[j]{
					son[j],son[i]=son[i],son[j]
				}


			}
		}
		
	}()

	
	return son
}

func main() {
	var son, n int

	fmt.Print("Slice count: ")
	fmt.Scan(&n)

	s := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Son:")
		fmt.Scan(&son)
		s[i] = son
	}

	fmt.Println("sort:", Sort(s))
}
