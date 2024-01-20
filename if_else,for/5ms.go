//2-masala
package main
import ("fmt")
func main() {
	var a int
	fmt.Print("yil=")
	fmt.Scan(&a)
  if (a%4==0&&a%100==0&&a%400==0) {
    fmt.Println("Kabisa yili")
  }  else { 
    fmt.Println("Kabisa yili emas")
  }
}