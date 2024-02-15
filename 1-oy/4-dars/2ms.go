package main

import "fmt"

func main() {
	var a,b,c int
	fmt.Print("1-tomon= ")
	fmt.Scan(&a)
	fmt.Print("2-tomon= ")
	fmt.Scan(&b)
	fmt.Print("3-tomon= ")
	fmt.Scan(&c)
	k := (a+b>c)&&(b+c>a)&&(a+c>b)
	fmt.Println("uchburchak bo'la oladi",k)}