package main

import "fmt"

func plus(num int,num1 int){
if num<num1{
	for i := num1; i <= num; i++{
	fmt.Print(i)}
}else{ 
	for i := num1; i >= num; i--{
	fmt.Print(i)}
 }
 }


func main() {
 var num, num1 int
 fmt.Print("Kiriting:")
 fmt.Scan(&num)
 fmt.Print("Kiriting:")
 fmt.Scan(&num1)
 plus(num,num1)
}