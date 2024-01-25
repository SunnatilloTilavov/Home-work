//1_misol
//
// map
//array,slice,map
package main

import (
	"fmt"
)

func main() {
	var a,b,c,g,f int
	numbers:=[10]int{0,-1,2,-3,4,-5,-6,0,7,1}
	for _ ,number:=range numbers{
		if number>0{
		a+=1
	    }
		if number<0{
			b+=1
		}
		if number==0{
			c+=1
		}
		if number%2==0&&!(number==0){
			f+=1
		}
		if (number%2==1)||(number%2==-1){
			g+=1
		}

	}
	fmt.Println("mustab=",a," manfiy=",b," nol=",c," juft=",f,"toqson",g)

}

