// map
//array,slice,map
package main

import "fmt"

func main() {
	numbers:=[]int(1,2,3,4,5,1,2,3)
	var newarr=make([]int,0,10)
	newmap:=make(map[int]bool)

	for _ ,number:=range numbers{
		if !newmap[number]
		newarr=append(newarr,number)
		newmap[number]=true
	}
	fmt.Println("newarr",newarr)

}