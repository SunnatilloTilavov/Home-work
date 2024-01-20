
package main
import ("fmt")

func main() {
	var i int
	fmt.Print("oy= ")
	fmt.Scan(&i)
    switch i { 
	case 1:
		fmt.Println("11")
	case 2:
		fmt.Println("22 ")
	case 3:
		fmt.Println("31 ")
	case 4:
		fmt.Println("43")			
	case 5:
		fmt.Println("5qw ")
	case 6:
		fmt.Println("1dasd")
	case 7:
		fmt.Println("2caw ")
	case 8:
		fmt.Println("3cwad ")
	case 9:
		fmt.Println("4ca ")			
	case 10:
		fmt.Println("5ca ")
	case 11:
		fmt.Println("1wcd")
	case 12:
		fmt.Println("2 ")
	default:
		fmt.Println("xato")	
	}
}