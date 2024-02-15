package main

import (
	"fmt"
	"strings"
)
func main(){
	str:="Non-declaration Statement outside function body"
	a:=strings.ToLower(str)
	b:=strings.ToUpper(str)
	fmt.Println("tolower=",a)
	fmt.Println("tolower=",b)
}