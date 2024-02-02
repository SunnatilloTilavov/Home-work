package main
import "fmt"
func main(){
var a int 
var ptr *int
var abs **int
&a=42
ptr=&a
abs=&ptr
fmt.Println(a)
fmt.Println(&a)
fmt.Println(*ptr)
fmt.Println(ptr)
fmt.Println(&ptr)
fmt.Println(**abs)
fmt.Println(&abs)
}


if ptr==nill{
	fmt.Println("pointer is nil")
}

ptr:=new(int)
*ptr=42
fmt.Println(*ptr)
