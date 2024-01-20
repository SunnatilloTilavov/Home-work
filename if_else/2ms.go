
//test
package main
import ("fmt")

func main() {
	var a,b,d float32
	fmt.Print("1-son=")
	fmt.Scan(&a)
	fmt.Print("2-son=")
	fmt.Scan(&b)
	fmt.Print("3-son=")
	fmt.Scan(&d)
  if (a > b&&a>d) {
    fmt.Println("1-son katta")
  } else if a < b&&b>d {
		fmt.Println("2-son katta")
  } else { 
    fmt.Println("3-son katta")
  }
}