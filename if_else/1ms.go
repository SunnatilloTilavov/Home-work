//test
package main
import ("fmt")

func main() {
	var a,b float32
	fmt.Print("1-son=")
	fmt.Scan(&a)
	fmt.Print("2-son=")
	fmt.Scan(&b)
  if (a > b) {
    fmt.Println("1-son katta")
  } else if a==b {
		fmt.Println("Teng")
  } else {
    fmt.Println("2-son katta")
  }
}