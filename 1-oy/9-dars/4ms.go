package main

import (
	"fmt")
func main(){
	var n,son int
	fmt.Print("massiv nechta sondan iborat=")
	fmt.Scan(&n)
	slice:=make([]int,n)
	fmt.Println(slice)
	for i:=0;i<n;i++{
		fmt.Print("son=")
	    fmt.Scan(&son)
		slice[i]=son
	}
	if n==1&&son==9{
	slice2:=make([]int,n)
	slice2[0]=1
	slice2[1]=0
	fmt.Print("slice",slice2)
	}else{
	a:=slice[n-1]+1
	d:=slice[n-2]
	b:=a/10
	c:=a%10
	slice[n-1]=c
	slice[n-2]=d+b
	fmt.Print("slice",slice)
	}


}