package main
import ("fmt")

func main() {
	var a,b int
	fmt.Print("a=")
	fmt.Scan(&a)
	fmt.Print("b=")
	fmt.Scan(&b)
	i :=a
	for i<=b{
		fmt.Println(i)
		i++
		
	}
}