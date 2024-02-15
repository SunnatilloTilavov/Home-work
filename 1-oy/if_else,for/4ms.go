//1-masala
package main
import ("fmt")

func main() {
	var a int
	fmt.Print("1-son=")
	fmt.Scan(&a)
  if (a%3==0) {
    fmt.Println("Fizz")
  } else if (a%5==0){
		fmt.Println("Buzz")
  } else if (a/3==0&&a/5==0){
    fmt.Println("FizzBazz")
  }else{
    fmt.Println("ikkala songa ham bo'linmaydi") 
  }
}