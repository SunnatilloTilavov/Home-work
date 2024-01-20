
package main
import ("fmt")

func main() {
	var i int
	fmt.Print("Baho= ")
	fmt.Scan(&i)
    switch i { 
	case 1:
		fmt.Println("1 yomon")
	case 2:
		fmt.Println("2 qoniqarsiz")
	case 3:
		fmt.Println("3 qoniqarli")
	case 4:
		fmt.Println("4 yaxshi")			
	case 5:
		fmt.Println("5 alo")
	default:
		fmt.Println("xato")	
	}
}