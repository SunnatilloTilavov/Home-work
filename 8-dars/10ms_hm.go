//5-ms
package main
import "fmt"
func main(){
	var n ,son,a int
	fmt.Print("n=")
	fmt.Scan(&n)
	slice:=make([]int,n)
	fmt.Println(slice)
	for i:=0;i<n;i++{
		fmt.Print("son=")
	    fmt.Scan(&son)
		slice[i]=son
	}
	a=slice[n-1]
	slice[n-1]=slice[0]
	slice[0]=a
    fmt.Println(slice)
}

