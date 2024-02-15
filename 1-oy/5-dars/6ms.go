package main
import ("fmt")

func main() {
	var a string
	fmt.Print("Tilni kiriting go,c++,python,java= ")
	fmt.Scan(&a)
    switch a  { 
	case "go":
		fmt.Println("Hello,go developer")
	case "c++":
		fmt.Println("Hello,c++ developer")
	case "python":
		fmt.Println("Hello,python developer")
	case "java":
		fmt.Println("Hello,java developer")			
	default:
		fmt.Println("Bunday til yo'q",a)	
	}
}