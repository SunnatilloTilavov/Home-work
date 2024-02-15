
package main
import ("fmt")

func main() {
	var i int
	fmt.Print("Sonni kiriting= ")
	fmt.Scan(&i)
    switch i { 
	case 1:
		fmt.Println("1 dushanba")
	case 2:
		fmt.Println("2 seshanba")
	case 3:
		fmt.Println("3 chorshnba")
	case 4:
		fmt.Println("4 payshanba")			
	case 5:
		fmt.Println("5 juma")
	case 6:
		fmt.Println("6 shanba")		
	case 7:
		fmt.Println("yakshanba")
	default:
		fmt.Println("xato")	
	}
}