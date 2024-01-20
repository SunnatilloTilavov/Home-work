package main
import ("fmt")

func main() {
	var a,i float32
	fmt.Print("narxi=")
	fmt.Scan(&a)
	i=0.1
	for i<=1{
		fmt.Println(i,"kg= ",i*a)
		i+=0.1
		
	}
}