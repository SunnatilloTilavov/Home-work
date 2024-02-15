package main

import "fmt"
func kopaytma(a float64, b int) float64 {
	var s float64=1
	for i := 1 ; i <= b; i++{
		s*=a

	  
	   }
  //fmt.Println(s)
  return s
}

func main() {

 var a float64 
 var b int 
 fmt.Print("A=")
 fmt.Scan(&a)
 fmt.Print("B=")
 fmt.Scan(&b)
 fmt.Println(kopaytma(a, b))

}