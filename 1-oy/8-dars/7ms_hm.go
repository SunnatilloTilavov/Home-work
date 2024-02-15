//2_misol
package main

import (
	"fmt"
)

func main() {
	var a,b,c,g,f  int
	var num int
	numbers:=make([]int,0,10)
	for i:=1;i<=10;i++{
		fmt.Print(i,"-son=")
	    fmt.Scan(&num)
		numbers=append(numbers,num)


	}


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

