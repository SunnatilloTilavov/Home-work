package main

import "fmt"
func factarial(a  int) int {
	s:=1
	for i := 1 ; i <=a; i++{
		s*=i

	  
	   }
  //fmt.Println(s)
  return s
}

func main() {
 var a int  
 fmt.Print("Son=")
 fmt.Scan(&a)
 fmt.Println(factarial(a))

}