package main
import ("fmt")

func main() {
	var a,b int
	fmt.Print("a=")
	fmt.Scan(&a)
	fmt.Print("b=")
	fmt.Scan(&b)
	c:=b
	i :=a
	for i<b {
		fmt.Println(c-1)
		c--
		i++
		
	}
}