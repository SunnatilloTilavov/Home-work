package main

import "fmt"

func kattason(a, b, c int) {

 if a + b==c {
  fmt.Println("1")
 } else {
	fmt.Println("0")
 }


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