//4-ms
package main

import (
	"fmt")
func main(){
	var n ,a,son int
	fmt.Print("n=")
	fmt.Scan(&n)
	slice:=make([]int,n)
	//fmt.Println(slice)
	for i:=0;i<n;i++{
		fmt.Print("son=")
	    fmt.Scan(&son)
		slice[i]=son
	}
	maxnum:=slice[0]
	minnum:=slice[0]
	for _, num:=range slice{
		if num<minnum{
			minnum=num
		}
		if num>maxnum{
			maxnum=num
		}
	}
	for i:=0;i<n;i++{
		if slice[i]==minnum{
		a=i
	}
	
    }
	slice[a]=maxnum
   minnum2:=slice[0]
	for _, sum:=range slice{
		if sum<minnum2{
			minnum2=sum
		}
	}

	fmt.Println("minnum2=",minnum2)
	
}