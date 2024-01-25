package main

import "fmt"

func plus(num int)int{
sum :=0

 for i := 1; i <= num; i++ {
	//sum:=string(i)
	sum += i
 }
 //fmt.Println(sum)
 return sum
}

func main() {
 var num int
 fmt.Print("Kiriting:")
 fmt.Scan(&num)
 fmt.Println(plus(num))
}