
package main
import ("fmt")
func calculator(a ,b,c int) {
var min,max int
    if (a > b) {
		max=a
		min=b
	  } else {
		max=b 
		min=a
	  } 
	  if (c> max) {
		fmt.Println("eng katta son",c)
	  } else {
		fmt.Println("eng katta son",max)
	  } 
	  if(c> min) {
		fmt.Println("eng kichik  son",min)
	  } else {
		fmt.Println("eng kichik son",c)
	  } 
	  
}
func main() {
	var a,b,c int
	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)
	fmt.Print("c= ")
	fmt.Scan(&c)
	calculator(a,b,c)
   
   }