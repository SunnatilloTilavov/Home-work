//array,slice,map
package main

import "fmt"

func main() {
	var arr=[]int{}
	slice1 := make([]int, 5, 8)
	fmt.Println("slice1",slice1)

	arr1:=[3]int{1,2,3}
	slice:=[]int{1,2,3,4,5}
    slice=append(slice, 10)
	fmt.Println("arr1",slice)
	fmt.Println("arr1",len(arr1))
	fmt.Println("arr1",cap(arr1))
	fmt.Println("arr1",arr1)
	fmt.Println("p",arr1)
	arr =append(arr, 20)


	arr2:=new([]int)
	*arr2=make([]int,1,10)
	fmt.Println((*arr2)[0])

	//map
	//map1 :=make(map(int),int)
	//map1:=("monday")="first day"
   // fmt.Println(map1)


  //key:="mon"
  //map1[key]="firs day"
  
}