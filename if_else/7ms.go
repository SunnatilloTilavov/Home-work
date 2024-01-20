//4-masala
package main
import ("fmt")
func main() {
	var c int
	var b,a string
	fmt.Print("Ismingiz=")
	fmt.Scan(&a)
	fmt.Print("Familiyangiz=")
	fmt.Scan(&b)
	fmt.Print("yoshingiz=")
	fmt.Scan(&c)
  if (c%2==1) {
    fmt.Println("ismingiz",a)
  }  else { 
    fmt.Println("familyangiz",b)
  }
}
