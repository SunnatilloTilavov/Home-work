package main

import "fmt"

func main() {

 var a,b,c string

 fmt.Print("1-Harf = ")
 fmt.Scan(&a)
 fmt.Print("2-Harf = ")
 fmt.Scan(&b)
 fmt.Print("3-Harf = ")
 fmt.Scan(&c)
 k := a<b&&b<c
 fmt.Println("ketma-ketlik :",k)
}