package main

import (
 "fmt"
)

func isPrime(num int) bool {
 if num < 2 {
  return false
 }
 for i := 2; i <= num; i++ {
  if num%i == 0 {
   return false
  }
 }
 return true
}

func main() {
 fmt.Println("Soniy tub bo'luvchilari:")
 for i := 0; i <= 100; i++ {
  if isPrime(i) {
   fmt.Printf("%d ", i)
  }
 }
}