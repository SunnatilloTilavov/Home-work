package main
import "fmt"
func plus(num int)int{
sum :=1

 for i := 1; i <= num; i++ {
	//sum:=string(i)
	sum *= i
 }
 return sum
 //fmt.Println(sum)
}

func main() {
 var num int
 fmt.Print("Kiriting:")
 fmt.Scan(&num)
 fmt.Println(plus(num))
}