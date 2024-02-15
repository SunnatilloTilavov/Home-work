package main
import ("fmt")

func main() {
	var i int
	fmt.Print("son= ")
	fmt.Scan(&i)
    switch i { 
	case 3:
		fmt.Println("uchburchak ")
	case 4:
		fmt.Println("to'rt burchak")
	case 5:
		fmt.Println("beshburchak")
	case 6:
		fmt.Println("oltiburchak")			
	case 7:
		fmt.Println("yettiburchak")
	default:
		fmt.Println("xato")	
	}
}