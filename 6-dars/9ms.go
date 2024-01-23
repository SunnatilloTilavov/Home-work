package main

import "fmt"

func ass(a float64,b float64)float64{

  f := a/b
 return f
}

func main() {

 var a,b ,f float64
 fmt.Print("Kiriting:")
 fmt.Scan(&a)
 fmt.Print("Kiriting:")
 fmt.Scan(&b)
 fmt.Print(float64(f))
 ass(a,b)
}