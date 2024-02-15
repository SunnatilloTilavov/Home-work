package main 

import ( 
	"fmt"
) 
type Num struct { 
    a int
	b int
} 
func main(){
	Num1:=Num{a:5,b:6}
	// & adresni jo'natadi
    sumabd(&Num1)
    fmt.Print(Num1.a,Num1.b)
}
// * adresini oladi
func sumabd(Num1* Num){
	Num1.a+=Num1.b	
}
// packeg "log"
_,err:=fmt.Scan(&a)
if
log.Println("ifbifn",err)
