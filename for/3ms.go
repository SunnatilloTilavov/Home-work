package main
import ("fmt")

func main() {
	var a,d float32
	var b string
	fmt.Print("a=")
	fmt.Scan(&a)
	fmt.Print("amal=")
	fmt.Scan(&b)
	fmt.Print("b=")
	fmt.Scan(&d)
  if (b=="+") {
    fmt.Println("a+b=",a+d)
  } else if b=="-"{
		fmt.Println("a-b ",a-d)
  } else if b=="*" {
		fmt.Println("a*b",a*d)
  } else { 
    fmt.Println("a/b=",a/d)
  }
}