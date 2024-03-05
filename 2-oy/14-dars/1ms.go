package main

import (
	"fmt"
	"sync"
)

func main() {
	var a,b int
	// defer wg.Done()
	var wg sync.WaitGroup
	fmt.Println("son:")
	fmt.Scan(&a)
	fmt.Println("go rutina:")
	fmt.Scan(&b)
	wg.Add(b)
	s:=1

   for i:=1;i<=a;i++{
		 s*=i
	}

fmt.Println("natija",s)	
}

func factarial()

