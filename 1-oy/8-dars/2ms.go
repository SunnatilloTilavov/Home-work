package main
import "fmt"
func array(s [10]int)int {
	a:=s[1]+s[2]+s[3]
	return a


}

func main() {
s:=[10]int{1,2,3,4,5,6,7,8,9,10}
fmt.Println(array(s))
}
