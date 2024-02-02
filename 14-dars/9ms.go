package main

import (
	"fmt")
func main(){
	var n,son int
	fmt.Print("massiv nechta sondan iborat=")
	fmt.Scan(&n)
	slice:=make([]int,n)
	slice2:=make([]int,n)
	fmt.Println(slice)
	for i:=0;i<n;i++{
		fmt.Print("son=")
	    fmt.Scan(&son)
		slice[i]=son
	}
	b:=n
	for i:=0;i<n;i++{
		slice2[b-1]=slice[i]
		b-=1


	}
fmt.Print(slice2)

}