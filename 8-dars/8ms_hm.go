//3_misol

package main

import (
	"fmt"
)

func main() {
	var f int
	c:=1
	slice:=[]int{1,2,3,4,5,6,7,8,9,10}
	for i:=0;i<=9;i+=2{
		c*=slice[i]
	}
	for a:=1;a<=9;a+=2{
		f+=slice[a]

	}
	fmt.Println("juft o'rindagi sonlar ko'paytmasi=",c,"toq o'rinagi=",f)

}

