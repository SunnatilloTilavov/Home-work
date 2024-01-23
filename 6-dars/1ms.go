package main

import "fmt"

func abs(x string) {

 for i := len(x) - 1; i >= 0; i-- {
  fmt.Print(string(x[i]))

 }
}

func main() {

 var x string

 fmt.Print("Kiriting:")
 fmt.Scan(&x)

 abs(x)
}