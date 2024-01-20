
package main
import ("fmt")

func main() {
	var i int
	fmt.Print("Raqam= ")
	fmt.Scan(&i)
    switch i { 
	case 1,2,12:
		fmt.Println("qish")
	case 3,4,5:
		fmt.Println("bahor")
	case 6,7,8:
		fmt.Println("yoz")
	case 9,10,11:
		fmt.Println("kuz")			
	
	default:
		fmt.Println("xato")	
	}
}