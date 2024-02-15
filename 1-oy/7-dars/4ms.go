
package main
import ("fmt")
func calculator(a ,b int,i string) {
    switch i { 
	case "qo'shish":
		fmt.Println(a+b)
	case "ayirish":
		fmt.Println(a-b)
	case "bo'lish":
		fmt.Println(a/b)
	case "ko'paytirish":
		fmt.Println(a*b)			
	default:
		fmt.Println("xato")	
	}
}
func main() {
	var a,b int
	var i string
	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)
	fmt.Print("amal= ")
	fmt.Scan(&i)
	calculator(a,b,i)
   
   }