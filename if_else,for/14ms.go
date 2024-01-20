package main
import ("fmt")

func main() {
	var a,b float32
	fmt.Print("1-son=")
	fmt.Scan(&a)
	fmt.Print("2-son=")
	fmt.Scan(&b)
  if (a > b) {
    fmt.Printf("1-son katta %f  2-son kichik %f",a,b)
  } else if a==b {
		fmt.Println("Teng")
  } else {
    fmt.Printf("2-son katta %f 1-son kichik %f",b,a)
  }
}