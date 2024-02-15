package main
import "fmt"
func printValue(i interface{}){
if v,ok :=i.(string);ok{
	fmt.Println("String",v)
}else if v,ok :=i.(int);ok{
	fmt.Println("Int",v)
}else if v,ok :=i.(float32);ok{
	fmt.Println("float",v)
}else if v,ok :=i.(bool);ok{
	fmt.Println("Int",v)
}else {
	fmt.Print("not")
}
}
func main(){
	printValue(1)

}