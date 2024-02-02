package main
import "fmt"
func pointer(a,b *int){
*a,*b=*b,*a

}
func main(){
var a int =42
fmt.Println(&a)
var b int =32
fmt.Println(&b)

pointer(&a,&b)
fmt.Println(a,b)
fmt.Println(&a)
fmt.Println(&b)

}