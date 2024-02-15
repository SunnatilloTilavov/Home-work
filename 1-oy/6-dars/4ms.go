package main

import "fmt"

func plus(num int,num1 int){
if num<num1{
	for i := num; i <= num1; i++{
	fmt.Print(i)}
}else{ 
	for i := num; i >= num1; i--{
	fmt.Print(i)}
 }
 }


func main() {
 var num, num1 int
 fmt.Print("Kiriting:")
 fmt.Scan(&num)
 fmt.Print("Kiriting:")
 fmt.Scan(&num1)
 //fmt.Println(plus(num,num1))
 plus(num,num1)
}