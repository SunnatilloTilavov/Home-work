package main

import "fmt"
func kattason(a float64, b int) {
	var s float64=1
	for i := 1 ; i <= b; i++{
		s*=a

	  
	   }
  fmt.Println(s)
}

func main() {

 var a float64
 var b int
 fmt.Print("A=")
 fmt.Scan(&a)
 fmt.Print("B=")
 fmt.Scan(&b)
 kattason(a,b)

}