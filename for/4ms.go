package main

import "fmt"

func main() {
 var input string

 fmt.Print("Kiruvchi qiymatni kiriting: ")
 fmt.Scan(&input)

 for i := 0; i < len(input); i++ {
  combination := input[i:] + input[:i]
  fmt.Println(combination)
 }
}