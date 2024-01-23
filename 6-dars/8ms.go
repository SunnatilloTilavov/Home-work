package main

import "fmt"

func son(a float64,b float64){

  f:=a/b
  fmt.Println(f)
}

func main() {

 var a, b  float64

 fmt.Print("A=")
 fmt.Scan(&a)
 fmt.Print("B=")
 fmt.Scan(&b)
 son(a,b)
}