package main
import ("fmt")
func main() {
	var c,a,b,d,e,f int
	fmt.Print("1-son")
	fmt.Scan(&a)
	fmt.Print("2-son")
	fmt.Scan(&b)
	fmt.Print("3-son")
	fmt.Scan(&c)
  if (a>0) {
    a=1
  }  else { 
    d=1
  }

  if (b>0) {
    b=1
  }  else { 
    e=1
  }
  if (c>0) {
    c=1
  }  else { 
    f=1
  }
  fmt.Println("musbat sonlar ",a+b+c)
  fmt.Println("manfiy sonlar ",d+e+f)
}