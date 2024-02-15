//3-masala
package main
import ("fmt")
func main() {
	var a int
	fmt.Print("Son=")
	fmt.Scan(&a)
  if (a%2==0) {
    fmt.Println("Kabisa yili",a*a)
  }  else { 
    fmt.Println("Kabisa yili emas",a*a*a)
  }
}