package main

import "fmt"

func abs(x string)string {
qw:=""

 for i := len(x) - 1; i >= 0; i-- {
  fmt.Print(string(x[i]))
  qw += string(x[i])
 }
 return qw
}

func main() {

 var x string

 fmt.Print("Kiriting:")
 fmt.Scan(&x)
 fmt.Print(string(x))
 abs(x)
}