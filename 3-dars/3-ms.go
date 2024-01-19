package main

import "fmt"

func main() {
 var (a,b,d float64 
 e,r,z float64=18.5,24.9,40
 g,h,q,w float64=25,29.9,30,39.9)
 fmt.Print("bo'y(sm)= ")
 fmt.Scan(&a)
 fmt.Print("vazn= ")
 fmt.Scan(&b)
 d=(b*100*100)/(a*a)
 fmt.Println("BMI= %f",d)
 j :=d<e&&e<r
 k := e<d&&d<r
 p := g<d&&d<h
 c := q<d&&d<w
 x := d>z&&e<r
 fmt.Println(" vazn yerishmaydi: %v",k)
 fmt.Println(" normal vazn: %v",p)
 fmt.Println(" biroz vazn ortiqcha: %v",c)
 fmt.Println(" xavfli: %v",x,j)
}