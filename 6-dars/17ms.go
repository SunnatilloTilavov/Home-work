package main

import "fmt"
func kattason(a int) {
	for i := 1 ; i <= a; i+=2{
		fmt.Print(i," ")
	  
	   }
}

func main() {

 var a int
 fmt.Print("N=")
 fmt.Scan(&a)
 kattason(a)
}