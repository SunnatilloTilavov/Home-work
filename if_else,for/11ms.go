package main
import ("fmt")
func main() {
	var c,a,b int
	fmt.Print("1-son")
	fmt.Scan(&a)
	fmt.Print("2-son")
	fmt.Scan(&b)
	fmt.Print("3-son")
	fmt.Scan(&c)
  if (a>0) {
    a=1
  }  else { 
    a=0
  }
  if (b>0) {
    b=1
  }  else { 
    b=0
  }
  if (c>0) {
    c=1
  }  else { 
    c=0
  }
  fmt.Println("musbat sonlar ",a+b+c)
}
