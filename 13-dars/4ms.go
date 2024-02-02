package main

import (
	"fmt"
	// "strconv"
)

func main(){
	text:="10"
	// b,err:=strconv.Atoi(text)
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println(b)

d:=fmt.Sprintf("%v",text)
fmt.Println(d)
}