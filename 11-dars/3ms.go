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
    sumab(&Num1)
    fmt.Print(Num1.a,Num1.b)
}
// * adresini oladi
func sumab(Num1* Num){
	Num1.a+=Num1.b	
}