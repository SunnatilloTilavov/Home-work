//array,slice,map
package main

import "fmt"

func main() {
	n:=0
	fmt.Scan(&n)
	soz:=make([]string,0,n)
	for i:=0 ; i<n; i++{
		tempNuumb :=" "
		fmt.Scan(&tempNuumb)
		soz=append(soz,tempNuumb)
	}
	fmt.Println(soz)

}