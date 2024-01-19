package main

import "fmt"

func main() {

 var a,b,d float64
 fmt.Print("bo'y(sm)= ")
 fmt.Scan(&a)
 fmt.Print("vazn= ")
 fmt.Scan(&b)
 d=(b*100*100)/(a*a)
 fmt.Println("BMI= ",d)
}