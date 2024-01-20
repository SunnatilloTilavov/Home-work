package main
import ("fmt")

func main() {
	var a int
	fmt.Print("narxi=")
	fmt.Scan(&a)
	
	i :=1
	for i<=10{
		fmt.Println(i,"kg=",i*a)
		i++
		
	}
}