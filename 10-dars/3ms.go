package main

import (
	"fmt"
	"strings"
)
func main(){
	str:="non-declaration statement outside function body"
	substr:="declaration"
	a:=strings.Index(str,substr)
	fmt.Println("Index=",a)
}