package main
import ("fmt"
   "time")
func main(){

	var a,b int
	fmt.Print("num=")
	fmt.Scan(&a)
	fmt.Print("time=")
	fmt.Scan(&b)
	now := time.Now()
	//fmt.Println(now)
	if b>3600 {
		fmt.Println("Vaqt juda katta ")
	}else{
		
	for i:=1;i<=a;i++{
		fmt.Println(i)
		time.Sleep(time.Duration(time.Second*time.Duration(b)))
	}
	}
	fmt.Println(time.Since(now))
}