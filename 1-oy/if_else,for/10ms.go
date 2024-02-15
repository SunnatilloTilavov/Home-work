package main
import ("fmt")
func main() {
	var c int
	fmt.Print("son=")
	fmt.Scan(&c)
  if (c>0) {
    fmt.Println("musbat",c+1)
  }  else if c==0{ 
    fmt.Println("manfiy",c+10)
  } else { 
    fmt.Println("manfiy",c+2)
  }
}
