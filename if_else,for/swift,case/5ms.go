
package main
import ("fmt")

func main() {
	var i,a,b float32
	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)
	fmt.Print("amal= ")
	fmt.Scan(&i)
    switch i { 
	case 1:
		fmt.Println(a+b)
	case 2:
		fmt.Println(a-b)
	case 3:
		fmt.Println(a/b)
	case 4:
		fmt.Println(a*b)			
	default:
		fmt.Println("xato")	
	}
}