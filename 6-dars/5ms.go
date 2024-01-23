package main

import "fmt"

func kattason(a, b, c int) {

 var min, max, min1, max1,f int

 if a > b {
  max = a
  min = b
 } else {
  max = b
  min = a
 }
 if max > c {
  max1 = max
 } else {
  max1 = c
 }
 if min < c {
  min1 = min
 } else {
  min1 = c
 }
  f=(min1+max1)/2
  fmt.Println("Oradagi son ",f)


}

func main() {

 var a, b, c int

 fmt.Print("A=")
 fmt.Scan(&a)
 fmt.Print("B=")
 fmt.Scan(&b)
 fmt.Print("C=")
 fmt.Scan(&c)

 kattason(a, b, c)

}