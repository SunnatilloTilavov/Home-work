package main

import "fmt"

func main() {

 var harf string

 fmt.Print("Harf = ")
 fmt.Scan(&harf)

 k := (("A" == harf)||("E" == harf)||("O" == harf)||("U" == harf)||("I" == harf))

 fmt.Printf("unli: %v\n",k)
}