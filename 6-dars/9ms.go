package main

import "fmt"

func ass(a float64,b float64)float64{

  f := a/b
 return f
}

func main() {

 var a,b  float64
 fmt.Print("A:")
 fmt.Scan(&a)
 fmt.Print("B:")
 fmt.Scan(&b)
 fmt.Println(ass(a,b))
}