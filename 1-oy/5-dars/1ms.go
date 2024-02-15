package main
import ("fmt")

func main() {
	var a,b int
	fmt.Print("n=")
	fmt.Scan(&a)
  if (a%2==0) {
    b=a
  } else {
    b=a-1
  }
	i :=b
	for i>=1{
		fmt.Println(i)
		i-=2
		
	}
}